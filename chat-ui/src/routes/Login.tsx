import React, {useState} from 'react';
import styled from 'styled-components';
import {useNavigate} from "react-router-dom";
import {PrimaryButton} from "../components/Buttons";
import {InputField} from "../components/FormFields";
import {TOKEN_STORAGE_KEY} from "../auth";
import {GREEN, LIGHT_BLACK, LIGHTEST_BLACK} from "../styles";

const Container = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: ${LIGHT_BLACK};
`;

const Box = styled.div`
  border-radius: 50px;
  background-color: ${LIGHTEST_BLACK};
  padding: 30px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

const ErrorMessage = styled.p`
  color: ${GREEN};
  font-size: 14px;
  margin-top: 10px;
  text-align: center;
`;


export default function Login(): JSX.Element {
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
        setPassword(event.target.value);
        setErrorMessage('');
    };

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const response = await fetch("http://localhost:8080/login", {
            method: "post",
            headers: {
                "Accept": "application/json"
            },
            body: JSON.stringify({
                password,
            })
        });
        if (!response.ok) {
            setErrorMessage("Something went wrong!");
        }
        const payload = await response.json();
        if (payload["token"]) {
            await localStorage.setItem(TOKEN_STORAGE_KEY, payload["token"])
            navigate("/chat")
        }
    }

    return (
        <Container>
            <form onSubmit={handleSubmit}>
                <Box>
                    <InputField
                        type="password"
                        id="password-input"
                        placeholder="password"
                        value={password}
                        onChange={handleChange}
                    />
                    {errorMessage && <ErrorMessage>{errorMessage}</ErrorMessage>}
                    <PrimaryButton disabled={!password} type="submit">Enter chat</PrimaryButton>
                </Box>
            </form>
        </Container>
    );
};