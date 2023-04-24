import Button from "@mui/material/Button";
import styled from "styled-components"

export const PrimaryButton = styled(Button)`
    && {
      color: #00FF00;
      border: none;
      
      &:hover:not(:disabled) {
        background-color: #000000;
      }
      
      &:disabled {
        color: #2e5c2e;
        cursor: not-allowed;
      }
  }
`;