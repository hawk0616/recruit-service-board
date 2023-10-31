import { useState, useEffect } from 'react';
import axios from 'axios';
import { Technology } from '../types/companyTechnology'
import { useError } from './useError'
import { useRouter } from 'next/router';

const useGetTechnologiesByCompanyId = () => {
  const [companyTechnologies, setCompanyTechnologies] = useState<Technology[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState(null);
  const { ErrorHandling } = useError();
  const { query } = useRouter();

  useEffect(() => {
    const fetchTechnologiesByCompanyId = async () => {
      if (!query.id) return;

      try {
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/${query.id}/company_technologies`, { withCredentials: true })
        setCompanyTechnologies(response.data)
        setLoading(false)
      } catch (err: any) {
        if (err.response && err.response.data) {
          ErrorHandling(err.response.data)
        } else if (err.response) {
          ErrorHandling(err.response)
        } else {
          ErrorHandling(err)
        }
        setError(err)
        setLoading(false)
      }
    };

    fetchTechnologiesByCompanyId();
  }, [query.id, ErrorHandling]);

  return { companyTechnologies, loading, error };
}

export default useGetTechnologiesByCompanyId
