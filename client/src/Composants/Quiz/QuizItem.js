// MenuItem.jsx
import { useContext } from 'react'
import { QuizContext, UserContext } from '../AppContext';
import axios from 'axios';


const QuizItem = ({ label }) => {
    const { handleLogout } = useContext(UserContext);
    const { quizType, setQuizData } = useContext(QuizContext);
    const handleClick = async () => {
        await axios({
            method: "POST",
            url: "http://localhost:8000/anime/answer",
            withCredentials: true,
            data: {
                answer: label, 
                quiz: quizType
            }
        })
        .then(res => {
            setQuizData(res.data);
        })
        .catch(err => {
            handleLogout();
        })
    };

    return (
        <button className="button" onClick={handleClick}>
            {label}
        </button>
    );
}

export default QuizItem;
