import React, {memo, useCallback} from "react";
import styled from "styled-components";
import CheckIcon from '@mui/icons-material/Check';
import {BLACK, DISABLED_GREEN, LIGHT_BLACK, LIGHT_GREEN, LIGHTEST_BLACK, RED} from "../styles";

const InternalMessage = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
  width: 75%;
  padding-right: 100px;
  padding-left: 10px;
  border-radius: 15px;
  margin: 10px;
`;

const OkMessage = styled(InternalMessage)``;

const BadMessage = styled(InternalMessage)`
    color: ${RED};
`;

const PendingMessage = styled(InternalMessage)`
    font-syle: italic;
    color: ${LIGHT_BLACK};
`;

const ExternalMessage = styled(InternalMessage)`
    justify-content: flex-end;
    padding-right: 10px;
    padding-left: 100px;
`;

const InternalSpeechBubble = styled.p`
    color: ${BLACK};
    background-color: ${LIGHT_GREEN};
    border-radius: 50px;
    padding: 3px 15px 5px 15px;
    margin: 0px;
`;

const ExternalSpeechBubble = styled(InternalSpeechBubble)`
    color: ${LIGHT_GREEN};
    background-color: ${BLACK};
`;

const Messages = styled.div`
    border-radius: 50px;
    background-color: ${LIGHTEST_BLACK};
    padding: 30px;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    align-items: center;
    height: 300px;
    width: 100%;
    max-width: 500px;
    margin: 20px;
    overflow: hidden;
`;

const SuccessMark = styled(CheckIcon)`
    && {
        transform: translate(-2px, 1px);
        font-size: 0.8rem;
    }
`;

export interface Message {
    text: string;
}

interface Props {
    messages: Message[];
    okMessages: { [key: string]: boolean };
    badMessages: { [key: string]: boolean };
    externalMessages: { [key: string]: boolean};
}

function MessageListComponent({ messages, okMessages, badMessages, externalMessages }: Props): JSX.Element {

    const getMessageInputText = useCallback((message: Message) => {
        if (message.text.length) {
            const messageParts = message.text.split("|");
            if (messageParts.length > 1) {
                return message.text.split("|")[1];
            }
        }

        return "";
    }, [messages]);

    const isExternalMessage = useCallback((message: Message) => {
        return Boolean(externalMessages[message.text]);
    }, [externalMessages]);

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

                if (isExternalMessage(message)) {
                    return (
                        <ExternalMessage data-testid="external-message" {...{props}}>
                            <ExternalSpeechBubble>{messageInput}</ExternalSpeechBubble>
                        </ExternalMessage>
                    )
                } else if (isOkMessage(message)) {
                    return (
                        <OkMessage data-testid="ok-message" {...{props}}>
                            <InternalSpeechBubble>
                                <SuccessMark />
                                {messageInput}
                            </InternalSpeechBubble>
                        </OkMessage>
                    )
                } else if (isBadMessage(message)) {
                    return (
                        <BadMessage data-testid="bad-message" {...{props}}>
                            <InternalSpeechBubble>{messageInput}</InternalSpeechBubble>
                        </BadMessage>
                    )
                }

                return (
                    <PendingMessage data-testid="pending-message" {...{props}}>
                        <InternalSpeechBubble>{messageInput}</InternalSpeechBubble>
                    </PendingMessage>
                );
            })}
        </Messages>
    );
}

const MessageList = memo(MessageListComponent);

export default MessageList;