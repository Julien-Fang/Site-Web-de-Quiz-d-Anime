import './App.css';
import { UserContext } from './Composants/AppContext';
import Routes from "./Composants/Routers/Routers";
import { useState, useEffect } from 'react';
import axios from 'axios';

require('bootstrap/dist/css/bootstrap.min.css')

function App() {
    const [userid, setUserid] = useState(null);
    const [pseudo, setPseudo] = useState("");

    useEffect(() => {
        const storedUser = localStorage.getItem("userid");
        if (storedUser) {
            setUserid(JSON.parse(storedUser));
        }
        if (userid) {
            axios({
                method: "GET",
                url: "http://localhost:8000/user",
                withCredentials: true,
            })
            .then((res) => {
                setPseudo(res.data.pseudo);
            })
            .catch((err) => {
                handleLogout();
            });
        }
    }, [userid]);

    const handleLogin = (userid) => {
        localStorage.setItem("userid", JSON.stringify(userid));
        setUserid(userid);
    };

    const handleLogout = () => {
        localStorage.removeItem("userid");
        setUserid(null);
    };
    
    return (
        <UserContext.Provider value={{handleLogin, handleLogout,
                                        userid,
                                        pseudo, setPseudo
        }}>
            <Routes />
        </UserContext.Provider>
    );
}

export default App;