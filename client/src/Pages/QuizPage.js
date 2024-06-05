import LeaderBoard from '../Composants/Quiz/LeaderBoard';
import Quiz from '../Composants/Quiz/Quiz';

const QuizPage = () => {
    return (
        <div className='quiz-page'>
            <div className='quiz-container row'>
                <Quiz/>
                <LeaderBoard/>
            </div>
        </div>
    );
};

export default QuizPage;