import { useState, useEffect } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';
import { CsrfToken } from '../types/auth';

export const useError = () => {
  const router = useRouter();
  const [message, setMessage] = useState<string | null>(null);

  const getCsrfToken = async () => {
    const { data } = await axios.get<CsrfToken>(
      `${process.env.NEXT_PUBLIC_API_URL}/csrf`
    );
    axios.defaults.headers.common['X-CSRF-TOKEN'] = data.csrf_token;
  };

  useEffect(() => {
    if (message) {
      const timer = setTimeout(() => setMessage(null), 5000); // 5秒後にメッセージをクリア
      return () => clearTimeout(timer);
    }
  }, [message]);

  const ErrorHandling = (msg: string) => {
    switch (msg) {
      case 'invalid csrf token':
        getCsrfToken();
        setMessage('CSRF token is invalid, please try again');
        break;
      case 'invalid or expired jwt':
        setMessage('access token expired, please login');
        router.push('/');
        break;
      default:
        setMessage(msg);
    }
  };

  return { ErrorHandling, flashMessage: message };
};
