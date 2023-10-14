import { Inter } from 'next/font/google'
import Link from 'next/link'
import { useEffect } from 'react'
import { useRouter } from 'next/router';
import { useAuth } from '@/hooks/useAuth'


const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const router = useRouter();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/auth');
    }
  }, [isAuthenticated, router]);

  return (
    <div>
      <Link href="/company" passHref>
        <div className="bg-gray-200 text-white font-bold py-2 px-4 rounded cursor-pointer">
          企業一覧へ
        </div>
      </Link>
    </div>
  );
}
