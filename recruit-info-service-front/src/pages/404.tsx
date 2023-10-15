import React from 'react';

const NotFoundPage = () => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h1 className="text-6xl font-bold text-red-500">404</h1>
        <p className="mt-4 text-xl font-semibold text-gray-700">ページが見つかりません。</p>
      </div>
    </div>
  );
};

export default NotFoundPage;