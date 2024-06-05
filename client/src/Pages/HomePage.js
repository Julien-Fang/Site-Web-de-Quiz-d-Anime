import { QuizContext } from '../Composants/AppContext';
import { useEffect, useState } from 'react';

import MenuPage from './MenuPage';
import QuizPage from './QuizPage';

const HomePage = () => {
    const [quizData, setQuizData] = useState(null);
    const [quizType, setQuizType] = useState("");
    const [num, setNum] = useState(0);

    useEffect(() => {
        if (quizData) {
            setNum(Number(quizData.number_question));
        }
    }, [quizData]);
    
    return (
        <QuizContext.Provider value={{quizData, setQuizData,
                                        quizType, setQuizType,
                                        num, setNum
        
        }}>
            {quizData ? <QuizPage /> : <MenuPage />}
        </QuizContext.Provider>
);
        
};

export default HomePage;