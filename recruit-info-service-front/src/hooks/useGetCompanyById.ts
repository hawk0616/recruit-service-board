import { useState, useEffect } from 'react';
import axios from 'axios';
import { Company } from '../types/company'
import { useError } from './useError'
import { useRouter } from 'next/router';

export const useGetCompanyById = () => {
  const [company, setCompany] = useState<Company>();
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState(null);
  const { switchErrorHandling } = useError();
  const { query } = useRouter();

  useEffect(() => {
    const fetchCompanyById = async () => {
      if (!query.id) return;

      try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/${query.id}`, { withCredentials: true })
        setCompany(response.data)
        setLoading(false)
      } catch (err: any) {
        if (err.response) {
          switchErrorHandling(err.response)
        } else {
          switchErrorHandling(err.response.data)
        }
        setError(err)
        setLoading(false)
      }
    };

    fetchCompanyById();
  }, [query.id]);

  return { company, loading, error };
};
