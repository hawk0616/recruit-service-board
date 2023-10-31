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
      className="p-4 border rounded-lg w-3/4"
      placeholder="企業名で検索"
    />
  );
};

export default SearchForm;
