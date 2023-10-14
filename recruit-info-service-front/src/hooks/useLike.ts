import { useState } from 'react';
import axios from 'axios';
import { useError } from './useError'

export const useLike = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState(null);
  const { switchErrorHandling } = useError();
  
  const createLike = async (companyId: number) => {
    setLoading(true);
    try {
      const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/companies/likes`, {companyId}, { withCredentials: true })
      setLoading(false);
    } catch (err: any) {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message)
      } else {
        switchErrorHandling(err.response.data)
      }
      setError(err)
      setLoading(false)
    }
  };

  const deleteLike = async (companyId: number) => {
    setLoading(true);
    try {
      const response = await axios.delete(`${process.env.NEXT_PUBLIC_API_URL}/companies/likes/${companyId}`, { withCredentials: true });
    } catch (err: any) {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message);
      } else {
        switchErrorHandling(err.response.data);
      }
      setError(err);
    }
    setLoading(false);
  };

  const checkLike = async (companyId: number): Promise<boolean | undefined> => {
    setLoading(true);
    try {
      const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/likes/${companyId}`, { withCredentials: true });
      setLoading(false);
      console.log(response.data.liked)
      return response.data.liked;
    } catch (err: any) {
      if (err.response.data.message) {
        switchErrorHandling(err.response.data.message);
      } else {
        switchErrorHandling(err.response.data);
      }
      setError(err);
      setLoading(false);
    }
  };
  
  return { 
    createLike,
    deleteLike,
    checkLike,
    loading,
    error
  };
};
