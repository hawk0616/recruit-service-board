import { useState } from 'react';
import axios from 'axios';
import { useError } from './useError';

export const useLike = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const { ErrorHandling } = useError();

  const handleAxiosError = (err: any) => {
    let errorMessage = 'An error occurred.';
    if (err.response && err.response.data.message) {
      errorMessage = err.response.data.message;
    } else if (err.response && err.response.data) {
      errorMessage = err.response.data;
    } else if (err.message) {
      errorMessage = err.message;
    }
    ErrorHandling(errorMessage);
    setError(errorMessage);
    setLoading(false);
  }

  const createLike = async (companyId: number) => {
    setLoading(true);
    try {
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/companies/likes`, { companyId }, { withCredentials: true });
      setLoading(false);
    } catch (err: any) {
      handleAxiosError(err);
    }
  };

  const deleteLike = async (companyId: number) => {
    setLoading(true);
    try {
      await axios.delete(`${process.env.NEXT_PUBLIC_API_URL}/companies/likes/${companyId}`, { withCredentials: true });
      setLoading(false);
    } catch (err: any) {
      handleAxiosError(err);
    }
  };

  const checkLike = async (companyId: number): Promise<boolean | undefined> => {
    setLoading(true);
    try {
      const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/likes/${companyId}`, { withCredentials: true });
      setLoading(false);
      return response.data.liked;
    } catch (err: any) {
      handleAxiosError(err);
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
