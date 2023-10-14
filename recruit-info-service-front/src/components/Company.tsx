import { useAuth } from '../hooks/useAuth'

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faArrowRightToBracket } from "@fortawesome/free-solid-svg-icons";

const Company = () => {
  const { logout } = useAuth();

  const handleLogout = async () => {
    await logout();
}

  return (
    <div>
      <FontAwesomeIcon 
        icon={faArrowRightToBracket}
        onClick={handleLogout}
        className="h-6 w-6 my-6 text-blue-500 cursor-pointer"
      />
    </div>
  )
}

export default Company
