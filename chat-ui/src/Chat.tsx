import React, {useEffect, useState} from 'react';
import MessageList, {Message} from "./components/MessageList";
import { v4 as uuid4 } from "uuid";
import styled from "styled-components";
import {PrimaryButton} from "./components/Buttons";
import {TOKEN_STORAGE_KEY} from "./auth";
import {useNavigate} from "react-router-dom";

interface SocketMessage {
    id: string;
    clientId: string;
    payload: string;
    originalPayload: string;
    type: string;
}


const Page = styled.div`
    width: 100vw;
    height: 100vh;
    background-color: #080808;
`;

const Header = styled.div`
    padding: 20px;
    width: 100vw;
    background-color: #000000;
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
`;

const Title = styled.h1`
    color: #00FF00;
`;

const Content = styled.div`
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
`;

export default function Chat(): JSX.Element {
    const [webSocket, setWebSocket] = useState<WebSocket | null>(null)
    const [messages, setMessages] = useState<Message[]>([]);
    const [queuedMessages, setQueuedMessages] = useState<Message[]>([]);
    const [newMessage, setNewMessage] = useState<string>('');
    const [okMessages, setOkMessages] = useState<{ [key: string]: boolean }>({ x: true });
    const [badMessages, setBadMessages] = useState<{ [key: string]: boolean }>({ x: true });

    const navigate = useNavigate()
    const handleLogout = () => {
        localStorage.removeItem(TOKEN_STORAGE_KEY);
        navigate("/login");
    }

    const handleOkMessage =(socketMessage: SocketMessage) => {
        setOkMessages(state => ({ ...state, [socketMessage.originalPayload]: true }));
        setBadMessages(state => ({ ...state, [socketMessage.originalPayload]: false }));
    };

    const handleBadMessage = (socketMessage: SocketMessage) => {
        setOkMessages(state => ({ ...state, [socketMessage.originalPayload]: false }));
        setBadMessages(state => ({ ...state, [socketMessage.originalPayload]: true }));
    };

    const handleTextMessage = (socketMessage: SocketMessage) => {
        setMessages(state => [...state, { text: socketMessage.payload }]);
        setTimeout(() => {
            handleOkMessage(socketMessage);
        }, 200);
    };

    const handleIncomingMessage = (messageText: string) => {
        const socketMessage: SocketMessage = JSON.parse(messageText);
        if (socketMessage.type == "TEXT") {
            handleTextMessage(socketMessage);
        } else if (socketMessage.type == "SUCCESS") {
            // Golang is too fast! Only doing this to highlight the "pending state" for a message
            setTimeout(() => {
                handleOkMessage(socketMessage);
            }, 200);
        } else if (socketMessage.type == "ERROR") {
            handleBadMessage(socketMessage);
        }
    };

    const handleCloseConnection = () => {
        setWebSocket(null);
    };

    useEffect(() => {
        if (!webSocket || webSocket.CLOSED) {
            const ws: WebSocket = new WebSocket('ws://localhost:8080/chat')
            ws.onopen = () => setWebSocket(ws);
            ws.onmessage = (event: MessageEvent) => handleIncomingMessage(event.data);
            ws.onclose = () => handleCloseConnection();
        }
        return () => {
            if (webSocket) {
                webSocket.close();
                setWebSocket(null);
            }
        };
    }, []);

    useEffect(() => {
        if (webSocket) {
            queuedMessages.forEach((queuedMessage) => {
                webSocket.send(queuedMessage.text);
            });
            setMessages(state => [...state, ...queuedMessages]);
            setQueuedMessages([]);
        }
        // By including webSocket in our dependencies, we ensure that
        // unsent messages get processed when new connections are made.
    }, [queuedMessages, webSocket]);


    // Callbacks
    function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        setNewMessage(event.target.value);
    }

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        if (newMessage) {
            setQueuedMessages(state => [...state, { text: `${uuid4()}|${newMessage}` }]);
            setNewMessage('');
        }
    };

    return (
        <Page>
            <Header>
                <Title>World Chat Go</Title>
                <PrimaryButton onClick={handleLogout}>Logout</PrimaryButton>
            </Header>
            <Content>
                <MessageList messages={messages} okMessages={okMessages} badMessages={badMessages}/>
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        value={newMessage}
                        onChange={handleChange}
                        placeholder="Type a message..."
                    />
                    <button type="submit">Send</button>
                </form>
            </Content>
        </Page>
    );
}