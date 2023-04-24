import React, {useEffect, useState} from 'react';
import MessageList, {Message} from "../components/MessageList";
import { v4 as uuid4 } from "uuid";
import styled from "styled-components";
import {PrimaryButton} from "../components/Buttons";
import {TOKEN_STORAGE_KEY} from "../auth";
import {useNavigate} from "react-router-dom";
import {ChatField} from "../components/FormFields";
import {BLACK, GREEN, LIGHT_BLACK} from "../styles";

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
    background-color: ${LIGHT_BLACK};
    overflow: hidden;
`;

const Header = styled.div`
    padding: 0 20px 0 20px;
    background-color: ${BLACK};
    display: flex;
    justify-content: space-between;
`;

const Title = styled.h1`
    color: ${GREEN};
`;

const Content = styled.div`
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
`;

const Form = styled.form`
    display: flex;
`;

export default function Chat(): JSX.Element {
    const [webSocket, setWebSocket] = useState<WebSocket | null>(null)
    const [messages, setMessages] = useState<Message[]>([]);
    const [queuedMessages, setQueuedMessages] = useState<Message[]>([]);
    const [newMessage, setNewMessage] = useState<string>('');
    const [okMessages, setOkMessages] = useState<{ [key: string]: boolean }>({ x: true });
    const [badMessages, setBadMessages] = useState<{ [key: string]: boolean }>({ x: true });
    const [externalMessages, setExternalMessages] = useState<{ [key: string]: boolean }>({ x: true });

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

    const handleExternalMessage = (socketMessage: SocketMessage) => {
        setMessages(state => [...state, { text: socketMessage.payload }]);
        setExternalMessages(state => ({ ...state, [socketMessage.payload]: true}))
        setTimeout(() => {
            handleOkMessage(socketMessage);
        }, 200);
    };

    const handleIncomingMessage = (messageText: string) => {
        const socketMessage: SocketMessage = JSON.parse(messageText);
        if (socketMessage.type == "TEXT") {
            handleExternalMessage(socketMessage);
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
    }

    return (
        <Page>
            <Header>
                <Title>World Chat Go</Title>
                <PrimaryButton onClick={handleLogout}>Logout</PrimaryButton>
            </Header>
            <Content>
                <MessageList
                    messages={messages}
                    okMessages={okMessages}
                    badMessages={badMessages}
                    externalMessages={externalMessages}
                />
                <Form onSubmit={handleSubmit}>
                    <ChatField
                        type="text"
                        value={newMessage}
                        onChange={handleChange}
                        placeholder="Type a message..."
                    />
                    <PrimaryButton type="submit">Send</PrimaryButton>
                </Form>
            </Content>
        </Page>
    );
}