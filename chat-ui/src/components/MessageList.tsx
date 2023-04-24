import React, {memo, useCallback} from "react";
import styled from "styled-components";

const OkMessage = styled.div``;
const BadMessage = styled.div`
    color: red;
`;
const PendingMessage = styled.div`
    font-style: italic;
    color: grey;
`;

const Messages = styled.div`
  border-radius: 50px;
  background-color: #101010;
  color: #00FF00;
  padding: 30px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 50%;
  width: 50%;
`;

export interface Message {
    text: string;
}

interface Props {
    messages: Message[];
    okMessages: { [key: string]: boolean };
    badMessages: { [key: string]: boolean };
}

function MessageListComponent({ messages, okMessages, badMessages }: Props): JSX.Element {

    const getMessageInputText = useCallback((message: Message) => {
        if (message.text.length) {
            const messageParts = message.text.split("|");
            if (messageParts.length > 1) {
                return message.text.split("|")[1];
            }
        }

        return "";
    }, [messages]);

    const isOkMessage = useCallback((message: Message) => {
        return Boolean(okMessages[message.text]);
    }, [okMessages]);

    const isBadMessage = useCallback((message: Message) => {
        return Boolean(badMessages[message.text]);
    }, [badMessages]);

    return (
        <Messages>
            {messages.map((message: Message, index: number) => {
                const props = { key: message.text };
                const messageInput = getMessageInputText(message);
                if (isOkMessage(message)) {
                    return  <OkMessage {...{props}}>{messageInput}</OkMessage>
                } else if (isBadMessage(message)) {
                    return  <BadMessage {...{props}}>{messageInput}</BadMessage>
                }
                return  <PendingMessage {...{props}}>{messageInput}</PendingMessage>
            })}
        </Messages>
    );
}

const MessageList = memo(MessageListComponent);

export default MessageList;