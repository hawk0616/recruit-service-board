import { useState } from 'react';
import LoginForm from '@/components/user/LoginForm';
import SignupForm from '@/components/user/SignupForm';

const Auth = () => {
  const [isLogin, setIsLogin] = useState<boolean>(true);

  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <h2 className="my-6">{isLogin ? 'ログイン' : '新規登録'}</h2>
      {isLogin ? (
        <LoginForm switchToSignup={() => setIsLogin(false)} />
      ) : (
        <SignupForm switchToLogin={() => setIsLogin(true)} />
      )}
    </div>
  );
}

export default Auth;