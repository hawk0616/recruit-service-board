import { useState, useEffect } from 'react';
import axios from 'axios';
import { Company } from '../types/company'
import { useError } from './useError'



export const useGetCompanies = () => {
  const [companies, setCompanies] = useState<Company[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState(null);
  const { switchErrorHandling } = useError()

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies`, { withCredentials: true })
        setCompanies(response.data)
        setLoading(false)
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

    fetchTasks();
  }, []);

  return { companies, loading, error };
};
