import { useState, useEffect } from 'react';
import axios from 'axios';
import { Technology } from '../types/companyTechnology'
import { useError } from './useError'
import { useRouter } from 'next/router';

const useGetTechnologiesByCompanyId = () => {
  const [companyTechnologies, setCompanyTechnologies] = useState<Technology[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState(null);
  const { switchErrorHandling } = useError();
  const { query } = useRouter();

  useEffect(() => {
    const fetchTechnologiesByCompanyId = async () => {
      if (!query.id) return;

      try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/${query.id}/company_technologies`, { withCredentials: true })
        setCompanyTechnologies(response.data)
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

    fetchTechnologiesByCompanyId();
  }, [query.id]);

  return { companyTechnologies, loading, error };
}

export default useGetTechnologiesByCompanyId
