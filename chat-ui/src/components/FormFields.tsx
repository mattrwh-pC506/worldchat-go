import Input from '@mui/material/Input';
import styled from "styled-components";

export const InputField = styled(Input)`
    && {
      color: #00FF00;
      background-color: transparent;
      border: none;
      border-bottom: 2px solid #00FF00;
      font-size: 16px;
      padding: 10px;
      margin-bottom: 20px;
      width: 100%;
      outline: none;
    
      &::placeholder {
        color: #6BFF7C;
        font-weight: bold;
      }
    
      &:focus {
        outline: none;
        border-bottom: 2px solid #6BFF7C;
      }
  }
`;