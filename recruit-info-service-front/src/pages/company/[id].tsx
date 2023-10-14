import Link from 'next/link';
import { useGetCompanyById } from '@/hooks/useGetCompanyById';

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowLeft } from "@fortawesome/free-solid-svg-icons";

import BorderLine from '../../components/BorderLine';
import LoadingSpinner from '@/components/LoadingSpinner';
import TechnologyTags from '@/components/TechnologyTags';
import LikeButton from '@/components/LikeButton';

const CompanyDetailPage = () => {
  const { company, loading } = useGetCompanyById();

  console.log(company)

  if (loading) {
    return <LoadingSpinner />;
  }

  return (
    <div className="bg-gray-200 min-h-screen flex flex-col items-start justify-start p-6 relative">
      <Link href="/company" passHref>
        <FontAwesomeIcon icon={faArrowLeft} size="2x" className="absolute top-4 left-4 text-gray-600 hover:underline" />
      </Link>

      <div className="m-8 flex items-center">
        <h1 className="text-gray-800 text-4xl font-bold tracking-wide">{company?.name}</h1>
        <LikeButton companyId={company?.id} className="ml-4" />
      </div>

      <div className="m-8">
        <h2 className="text-gray-600 text-xl font-semibold">企業情報</h2>
        <BorderLine />
        <div className="relative mb-4 w-full max-w-7xl">
          <h2 className="text-gray-600 font-semibold mb-2 absolute top-0 left-0 px-2">事業内容</h2>
          <p className="text-gray-600 pl-28">{company?.description || "-"}</p>
        </div>
        <div className="relative mb-4 w-full max-w-7xl">
          <h2 className="text-gray-600 font-semibold mb-2 absolute top-0 left-0 px-2">OpenSalary</h2>
          <a href={company?.open_salary} target="_blank" rel="noopener noreferrer" className="text-gray-500 pl-28 underline hover:text-gray-700">
            {company?.open_salary || "-"}
          </a>
        </div>
        <div className="relative mb-4 w-full max-w-7xl">
          <h2 className="text-gray-600 font-semibold mb-2 absolute top-0 left-0 px-2">所在地</h2>
          <p className="text-gray-600 pl-28">{company?.address || "-"}</p>
        </div>
      </div>

      <div className="m-8">
        <h2 className="text-gray-600 text-xl font-semibold">技術情報</h2>
        <BorderLine />
        <TechnologyTags />
      </div>
    </div>

  );
};

export default CompanyDetailPage;
