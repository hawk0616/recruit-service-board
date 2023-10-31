import { useState, useCallback } from 'react';
import axios from 'axios';
import { useError } from './useError'

export const useComment = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState(null);
  const { ErrorHandling } = useError();
  
  const createComment = async (companyId: number, content: string) => {
    setLoading(true);
    try {
      const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/companies/comments`, {
        companyId,
        content
      }, { withCredentials: true })
      setLoading(false);
      return response.data;
    } catch (err: any) {
      if (err.response.data.message) {
        ErrorHandling(err.response.data.message)
      } else {
        ErrorHandling(err.response.data)
      }
      setError(err)
      setLoading(false)
    }
  };

  const deleteComment = async (companyId: number) => {
    setLoading(true);
    try {
      const response = await axios.delete(`${process.env.NEXT_PUBLIC_API_URL}/companies/comments/${companyId}`, { withCredentials: true });
    } catch (err: any) {
      if (err.response.data.message) {
        ErrorHandling(err.response.data.message);
      } else {
        ErrorHandling(err.response.data);
      }
      setError(err);
    }
    setLoading(false);
  };

  const getCommentsByCompanyId = useCallback(async (companyId: number) => {
    setLoading(true);
    try {
      const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/comments/${companyId}`, );
      setLoading(false);
      return response.data;
    } catch (err: any) {
      if (err.response && err.response.data.message) {
        ErrorHandling(err.response.data.message);
      } else if (err.response && err.response.data) {
        ErrorHandling(err.response.data);
      } else {
        ErrorHandling(err.message);
      }
      setError(err);
      setLoading(false);
    }
  }, [ErrorHandling]);

  return { 
    createComment,
    deleteComment,
    getCommentsByCompanyId,
    loading,
    error
  };
};