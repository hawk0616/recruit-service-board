import { useState, FormEvent } from 'react'
import { useAuth } from '../hooks/useAuth'
import { useError } from '@/hooks/useError';

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleCheck, faArrowsRotate } from "@fortawesome/free-solid-svg-icons";

const Auth = () => {
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [isLogin, setIsLogin] = useState<boolean>(true);
  const { login, signup } = useAuth();
  const { switchErrorHandling } = useError();

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      if (isLogin) {
        await login({ name: '', email, password: password });
      } else {
        await signup({ name: name, email, password: password });
        await login({ name: '', email, password: password });
      }
    } catch (err: any) {
      switchErrorHandling(err.response?.data?.message || err.message || "something went wrong");
    }
  };
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center">
        <FontAwesomeIcon icon={faCircleCheck} className="h-8 w-8 mr-2 text-blue-500" />
        <span className="text-center text-3xl font-extrabold">
          Recruit Info Service
        </span>
      </div>
      <h2 className="my-6">{isLogin ? 'ログイン' : '新規登録'}</h2>
      <form onSubmit={submitAuthHandler}>
        <div>
          <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="name"
            type="name"
            autoFocus
            placeholder="ユーザー名"
            onChange={(e) => setName(e.target.value)}
            value={name}
          />
        </div>
        <div>
          <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="email"
            type="email"
            autoFocus
            placeholder="メールアドレス"
            onChange={(e) => setEmail(e.target.value)}
            value={email}
          />
        </div>
        <div>
          <input
            className="mb-3 px-3 text-sm py-2 border border-gray-300"
            name="password"
            type="password"
            placeholder="パスワード"
            onChange={(e) => setPassword(e.target.value)}
            value={password}
          />
        </div>
        <div className="flex justify-center my-2">
          <button
            className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
            disabled={!name || !email || !password}
            type="submit"
          >
            {isLogin ? 'ログイン' : '新規登録'}
          </button>
        </div>
      </form>
      <FontAwesomeIcon
        icon={faArrowsRotate}
        onClick={() => setIsLogin(!isLogin)}
        className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
      />
    </div>
  )
}

export default Auth;
