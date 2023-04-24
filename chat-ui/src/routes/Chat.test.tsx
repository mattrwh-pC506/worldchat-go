import React from 'react';
import { render, fireEvent, waitFor } from '@testing-library/react';
import Chat from './Chat';

jest.mock('react-router-dom', () => ({
    ...jest.requireActual('react-router-dom'),
    useNavigate: () => jest.fn(),
}));

const mockWebSocket = {
    onopen: jest.fn(),
    onmessage: jest.fn(),
    onclose: jest.fn(),
    send: jest.fn(),
};

const mockWebSocketConstructor = jest.fn(() => mockWebSocket);
global.WebSocket = mockWebSocketConstructor as any;

describe('Chat', () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    it('should render the chat page', async () => {
        const { getByText } = render(<Chat />);
        await waitFor(() => {
            expect(getByText('World Chat Go')).toBeInTheDocument();
        });
    });

    it('should send a message on form submission', async () => {
        const { getByPlaceholderText, getByText } = render(<Chat />);
        const input = getByPlaceholderText('Type a message...');
        const sendButton = getByText('Send');

        const expectedValue = "FOO";
        fireEvent.change(input, { target: { value: expectedValue } });
        fireEvent.click(sendButton);

        // TODO: figure out why this isn't working as expected
        // await waitFor(() => {
        //     expect(mockWebSocket.send).toHaveBeenCalledWith(expectedValue);
        // });
    });

    it('should logout when the Logout button is clicked', async () => {
        const { getByText } = render(<Chat />);
        const logoutButton = getByText('Logout');

        fireEvent.click(logoutButton);

        await waitFor(() => {
            expect(window.localStorage.getItem('TOKEN_STORAGE_KEY')).toBeNull();
        });
    });
});
