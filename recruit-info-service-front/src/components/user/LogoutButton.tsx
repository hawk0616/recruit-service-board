import { useAuth } from "@/hooks/useAuth";

const LogoutButton = () => {
  const { logout } = useAuth();
  return (
    <button 
      onClick={logout}
      className="bg-black text-white px-4 py-2 rounded hover:bg-red-700 focus:outline-none transition"
    >
      ログアウト
    </button>
  );
};

export default LogoutButton;
