import { useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { UserContext } from '../AppContext';

const Recent = () => {
        const { handleLogout } = useContext(UserContext);
        const [recent, setRecent] = useState([]);
    
        useEffect(() => {
            axios({
                method : 'GET',
                url : `http://localhost:8000/anime/recent`,
                withCredentials : true
            })
            .then (res => {
                setRecent(res.data.slice(0, 5));
            })
            .catch(err => {
                handleLogout();
            })
        }, [])
    
        return (
            recent.length === 0 ? <p style={{color: 'red', fontSize:'25px'}}>Chargement...</p> :

            <div className="card-container">
                <div className="card-header">Derniers animés</div>
                <ul className="list-group">
                    {recent.map((anime, index) => (
                        <li key={index} className="list-item">
                            <span className="index">{index + 1}.</span>
                            <span className="name">{anime}</span>
                        </li>
                    ))}
                </ul>
            </div>
        );
}

export default Recent;