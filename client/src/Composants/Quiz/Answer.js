import React, { useEffect, useContext } from 'react';

import { QuizContext } from '../AppContext';
import QuizItem from './QuizItem';

const Answer = ({num}) => {
    const { quizData } = useContext(QuizContext);

    return (
        <div className="d-flex" style={{flexDirection: "column"}}>
                <div className="quiz-grid">
                    <QuizItem label={quizData.quizcollection[num].reponsepossible[0]} />
                    <QuizItem label={quizData.quizcollection[num].reponsepossible[1]} />
                    <QuizItem label={quizData.quizcollection[num].reponsepossible[2]} />
                    <QuizItem label={quizData.quizcollection[num].reponsepossible[3]} />
                </div>
        </div>
    );
}

export default Answer;