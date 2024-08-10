import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function AddTodo() {
    const [name, setName] = useState('');
    const [text, setText] = useState('');
    const [timeLimit, setTimeLimit] = useState('');
    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.post('http://localhost:8080/Todo/add', {
            Name: name,
            Text: text,
            TimeLimit: parseInt(timeLimit, 10)
        })
            .then(() => {
                navigate('/');
            })
            .catch(error => {
                console.error('Error adding todo:', error);
            });
    };

    return (
        <div>
            <h1>Add Todo</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    Name:
                    <input type="text" value={name} onChange={e => setName(e.target.value)} />
                </label>
                <label>
                    Text:
                    <input type="text" value={text} onChange={e => setText(e.target.value)} />
                </label>
                <label>
                    Time Limit (in days):
                    <input type="number" value={timeLimit} onChange={e => setTimeLimit(e.target.value)} />
                </label>
                <button type="submit">Add Todo</button>
            </form>
        </div>
    );
}

export default AddTodo;
