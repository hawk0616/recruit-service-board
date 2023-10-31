import { useState, FormEvent } from 'react';
import { useAuth } from '@/hooks/useAuth';
import { useError } from '@/hooks/useError';

const SignupForm = ({ switchToLogin }: { switchToLogin: () => void }) => {
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const { signup, login } = useAuth();
  const { ErrorHandling } = useError();

  const submitSignupHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      await signup({ name, email, password });
      await login({ name, email, password });
    } catch (err: any) {
      ErrorHandling(err.response?.data?.message || err.message || "something went wrong");
    }
  };

  return (
    <form onSubmit={submitSignupHandler} className="w-full max-w-2xl">
      <div>
        <input
          className="mb-4 px-4 text-lg py-3 border border-gray-300 w-full"
          name="name"
          type="text"
          placeholder="ユーザー名"
          onChange={(e) => setName(e.target.value)}
          value={name}
        />
      </div>
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
          disabled={!name || !email || !password}
          type="submit"
        >
          新規登録
        </button>
        <span className="text-indigo-600 cursor-pointer" onClick={switchToLogin}>
          ログインはこちら
        </span>
      </div>
    </form>
  );
}

export default SignupForm;
