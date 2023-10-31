import { useState, FormEvent } from 'react';
import { useAuth } from '@/hooks/useAuth';
import { useError } from '@/hooks/useError';

const LoginForm = ({ switchToSignup }: { switchToSignup: () => void }) => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const { login } = useAuth();
  const { ErrorHandling } = useError();

  const submitLoginHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      await login({
        email, password,
        name: null
      });
    } catch (err: any) {
      ErrorHandling(err.response?.data?.message || err.message || "something went wrong");
    }
  };

  return (
    <form onSubmit={submitLoginHandler} className="w-full max-w-2xl">
      <div>
        <input
          className="mb-4 px-4 text-lg py-3 border border-gray-300 w-full"
          name="email"
          type="email"
          placeholder="メールアドレス"
          onChange={(e) => setEmail(e.target.value)}
          value={email}
        />
      </div>
      <div>
        <input
          className="mb-4 px-4 text-lg py-3 border border-gray-300 w-full"
          name="password"
          type="password"
          placeholder="パスワード"
          onChange={(e) => setPassword(e.target.value)}
          value={password}
        />
      </div>
      <div className="flex justify-between items-center">
        <button
          className="py-3 px-6 text-lg rounded text-white bg-indigo-600"
          disabled={!email || !password}
          type="submit"
        >
          ログイン
        </button>
        <span className="text-indigo-600 cursor-pointer" onClick={switchToSignup}>
          新規登録はこちら
        </span>
      </div>
    </form>
  );
}

export default LoginForm;