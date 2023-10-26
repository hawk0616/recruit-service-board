const SearchButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <button onClick={onClick} className="ml-4 p-2 bg-blue-500 text-white rounded hover:bg-blue-600">
      検索
    </button>
  );
};

export default SearchButton;