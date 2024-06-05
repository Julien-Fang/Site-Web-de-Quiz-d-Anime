import { useContext, useState, useEffect } from 'react';
import { QuizContext, UserContext } from '../AppContext';
import axios from 'axios';

const LeaderBoard = () => {
    const { quizType, setQuizData } = useContext(QuizContext);
    const { handleLogout } = useContext(UserContext);
    const [leaderboard, setLeaderboard] = useState([]);

    useEffect(() => {
        axios({
            method : 'GET',
            url : `http://localhost:8000/stats/${quizType}`,
            withCredentials : true
        })
        .then (res => {
            setLeaderboard(res.data);
        })
        .catch(err => {
            handleLogout();
        })
    }
    , [quizType, setQuizData])

    return (
        <div className="col-md-4 d-flex flex-column align-items-center justify-content-center">
            <div className="card-header text-light">LeaderBoard</div>
            <div className="card-container">
                <ul className="list-group">
                    <li key="header" className="list-item">
                        <span className="index-ld col-md-4">Joueur</span>
                        <span className="index-ld col-md-4">Parties lancées</span>
                        <span className="index-ld col-md-4">Score</span>
                    </li>
                </ul>
                <ul className="list-group">
                    {leaderboard.map((player) => (
                        <li key={player.idjoueur} className="list-item">
                            <span className="name-ld col-md-4">{player.idjoueur}</span>
                            <span className="name-ld col-md-4">{player.nbquizgeneral}</span>
                            <span className="name-ld col-md-4">{player.score}</span>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
}

export default LeaderBoard;