import Link from 'next/link';
import { useGetCompanies } from '../../hooks/useGetCompanies';
import { Company } from '../../types/company'
import LoadingSpinner from '../../components/LoadingSpinner';
import SearchForm from '@/components/search/SearchForm';
import SearchButton from '@/components/search/SearchButton';
import { useState } from 'react';
import useSearchCompany from '@/hooks/useSearchCompany';
import LogoutButton from '@/components/user/LogoutButton';

const CompanyPage = () => {
  const { companies, loading } = useGetCompanies();
  const [searchInput, setSearchInput] = useState<string>(''); 
  const [searchQuery, setSearchQuery] = useState<string>(''); 
  const { companies: searchedCompanies, loading: searchLoading, setShouldSearch } = useSearchCompany(searchQuery);

  if (loading || searchLoading) {
    return <LoadingSpinner />;
  }

  const handleSearch = () => {
    setSearchQuery(searchInput);
    setShouldSearch(true);
  };

  const displayCompanies = searchQuery ? searchedCompanies : companies;


  return (
    <div className="bg-gray-200 p-6 rounded-lg shadow-md">
      <div className="absolute top-4 right-4"><LogoutButton /></div>
      <h1 className="text-black text-5xl font-bold tracking-wide mb-12">企業一覧</h1>
      <div className="mb-36">
        <SearchForm onSearch={(query) => setSearchQuery(query)} />
        <SearchButton onClick={handleSearch} />
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        {displayCompanies.map((company: Company) => (
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
