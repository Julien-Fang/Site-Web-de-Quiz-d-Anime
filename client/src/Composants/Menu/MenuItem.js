// MenuItem.jsx
import { useContext } from 'react'
import { QuizContext, UserContext } from '../AppContext';

const MenuItem = ({ label }) => {
    const { setQuizType } = useContext(QuizContext);
    const { handleLogout } = useContext(UserContext);

    const handleClick = () => {
        try {
            setQuizType(label.toLowerCase() + 'Quiz');
        } catch (error) {
            handleLogout();
        }
    };

    return (
        <button className="button" onClick={handleClick}>
            {label}
        </button>
    );
}


export default MenuItem;
