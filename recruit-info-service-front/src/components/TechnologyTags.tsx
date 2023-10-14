import React from 'react';
import useGetTechnologiesByCompanyId from '@/hooks/useGetTechnologiesByCompanyId';
import LoadingSpinner from './LoadingSpinner';
import Error from './Error';

const TechnologyTags = ({  }) => {
  const { companyTechnologies, loading, error } = useGetTechnologiesByCompanyId();

  if (loading) return <LoadingSpinner />;
  if (error) return <Error message={error} />;

  if (!companyTechnologies || companyTechnologies.length === 0) {
    return <p>不明</p>;
  }

  return (
    <div className="flex flex-wrap gap-2">
      {companyTechnologies.map((technology, index) => (
        <span key={index} className="bg-green-400 text-white px-3 py-1 rounded-full text-sm">
          {technology.name}
        </span>
      ))}
    </div>
  );
};

export default TechnologyTags;
