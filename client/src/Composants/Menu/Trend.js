import { useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { UserContext } from '../AppContext';

const Trend = () => {
    const { handleLogout } = useContext(UserContext);
    const [trend, setTrend] = useState([]);

    useEffect(() => {
        axios({
            method : 'GET',
            url : `http://localhost:8000//anime/trend`,
            withCredentials : true
        })
        .then (res => {
            setTrend(res.data.slice(0, 5));
        }
        )
        .catch(err => {
            handleLogout();
        })
    }, [])

    return (
        trend.length === 0 ? <p style={{color: 'red', fontSize:'25px'}}>Chargement...</p> : 

        <div className="card-container">
            <div className="card-header">Tendances</div>
            <ul className="list-group">
                {trend.map((anime, index) => (
                    <li key={index} className="list-item">
                        <span className="index">{index + 1}.</span>
                        <span className="name">{anime}</span>
                    </li>
                ))}
            </ul>
        </div>


    
    );

}

export default Trend;