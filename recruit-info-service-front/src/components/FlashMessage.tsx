const FlashMessage = ({ message }: { message: string | null }) => {
  if (!message) return null;

  return (
    <div className="fixed top-0 left-0 w-full p-4 bg-red-600 text-white text-center">
      {message}
    </div>
  );
};

export default FlashMessage;
