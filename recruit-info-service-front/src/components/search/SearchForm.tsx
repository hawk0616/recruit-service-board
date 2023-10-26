import { useState } from 'react';

const SearchForm = ({ onSearch }: { onSearch: (query: string) => void }) => {
  const [searchQuery, setSearchQuery] = useState<string>('');

  return (
    <input
      type="text"
      value={searchQuery}
      onChange={(e) => {
        setSearchQuery(e.target.value);
        onSearch(e.target.value);
      }}
      className="p-2 border rounded w-2/3"
      placeholder="企業名で検索"
    />
  );
};

export default SearchForm;
