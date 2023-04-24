import React from 'react';
import {render, fireEvent, screen, waitFor} from '@testing-library/react';
import Login from './Login';
import {MemoryRouter} from "react-router-dom";

const mockNavigate = jest.fn();
jest.mock('react-router-dom', () => ({
    ...jest.requireActual('react-router-dom'),
    useNavigate: () => mockNavigate,
}));

describe('Login component', () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    it('renders the login form', () => {
        const { getByTestId } = render(<MemoryRouter><Login /></MemoryRouter>);

        expect(getByTestId('password-input')).toBeInTheDocument();
        expect(getByTestId('submit-login')).toBeInTheDocument();
    });

    it('disables the button when password is empty', () => {
        const { getByTestId } = render(<MemoryRouter><Login /></MemoryRouter>);

        expect(getByTestId('submit-login')).toBeDisabled();
    });

    it('enables the button when password is entered', () => {
        const { getByTestId } = render(<MemoryRouter><Login /></MemoryRouter>);

        const passwordInput = screen.getByTestId('password-input');
        fireEvent.change(passwordInput, {target: {value: 'password'}});
        expect(getByTestId('submit-login')).toBeEnabled();
    });

    it('displays an error message on failed login attempt', async () => {
        jest.spyOn(window, 'fetch').mockResolvedValueOnce({
            ok: false,
            json: async () => ({error: 'Invalid password'}),
        } as any);

        const { getByTestId, getByText } = render(<MemoryRouter><Login /></MemoryRouter>);

        const passwordInput = getByTestId('password-input');
        fireEvent.change(passwordInput, {target: {value: 'password'}});
        fireEvent.submit(screen.getByRole('button'));

        await waitFor(() => {
            expect(getByText('Something went wrong!')).toBeInTheDocument();
        });
    });

    it('redirects to chat page on successful login', async () => {
        const token = 'token123';
        jest.spyOn(window, 'fetch').mockResolvedValueOnce({
            ok: true,
            json: async () => ({token}),
        } as any);

        const { getByTestId } = render(<MemoryRouter><Login /></MemoryRouter>);

        const passwordInput = getByTestId('password-input');
        fireEvent.change(passwordInput, {target: {value: 'password'}});
        fireEvent.submit(getByTestId('submit-login'));

        await waitFor(() => {
            expect(mockNavigate).toHaveBeenCalledWith('/chat');
        });
    });
});
