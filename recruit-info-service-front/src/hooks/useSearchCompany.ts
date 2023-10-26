import { useState, useEffect } from 'react';
import axios from 'axios';
import { Company } from '../types/company';
import { useError } from './useError';

const useSearchCompany = (name: string) => {
  const [companies, setCompanies] = useState<Company[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [shouldSearch, setShouldSearch] = useState<boolean>(false);
  const { switchErrorHandling } = useError();

  useEffect(() => {
    if (!shouldSearch){
      setLoading(false)
      return;
    }

    const fetchData = async () => {
      try {
        console.log("name", name)
        const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/companies/search?name=${name}`, { withCredentials: true });
        setCompanies(response.data);
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

    fetchData();
    setShouldSearch(false);
  }, [name, shouldSearch]);

  return { companies, loading, error, setShouldSearch };
};

export default useSearchCompany;

