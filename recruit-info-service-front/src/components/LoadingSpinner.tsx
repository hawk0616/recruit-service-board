import React from 'react';

const LoadingSpinner = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-200">
      <div className="w-16 h-16 border-t-4 border-white rounded-full animate-spin"></div>
    </div>
  );
};

export default LoadingSpinner;