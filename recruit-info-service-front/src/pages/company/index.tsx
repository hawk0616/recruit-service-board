import Link from 'next/link';
import { useGetCompanies } from '../../hooks/useGetCompanies';
import { Company } from '../../types/company'
import LoadingSpinner from '../../components/LoadingSpinner';

const CompanyPage = () => {
  const { companies, loading } = useGetCompanies();

  if (loading) {
    return <LoadingSpinner />;
  }

  return (
    <div className="bg-gray-200 p-6 rounded-lg shadow-md">
      <h1 className="text-black text-5xl font-bold tracking-wide mb-12">企業一覧</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        {companies.map((company: Company) => (
          <Link key={company.id} href={`/company/${company.id}`} passHref>
            <div className="block bg-white p-4 rounded shadow hover:bg-gray-100 transition cursor-pointer">
              <h2 className="text-xl font-semibold mb-2">{company.name}</h2>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default CompanyPage;
