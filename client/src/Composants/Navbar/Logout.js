import axios from 'axios';
import cookie from 'js-cookie';
import { useContext } from 'react';
import { UserContext } from '../AppContext';
import { BoxArrowRight } from 'react-bootstrap-icons';

const Logout = () => {

    const { handleLogout } = useContext(UserContext);

    const removeCookie = (key) => {
        if (window !== "undefined") {
            cookie.remove(key, { expires: 1 });
        }
    };

    const logout = async () => {
        await axios({
            method: "DELETE",
            url: "http://localhost:8000/user/logout",
            withCredentials: true,
        })
        .then((res) => {
            removeCookie("refreshToken");
            removeCookie("accessToken");
            removeCookie("id");
            handleLogout();
        })
        .catch(err => {
            handleLogout();
        });
    }
    
    return (
        <div className="logout d-flex justify-content-center" onClick={logout}>
            <span style={{ marginRight: "5px" }}>Logout</span>
            <BoxArrowRight />
        </div>
    )
  
}

export default Logout;