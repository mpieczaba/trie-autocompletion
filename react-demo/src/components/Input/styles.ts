import styled from "styled-components";

export const InputWrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 2rem;
  margin: 1rem;
`;

export const StyledInput = styled.input`
  display: flex;
  width: 40vw;
  height: 3rem;
  background-color: #282c34;
  color: #f2f2f2;
  font-size: 1.2rem;
  border: none;
  resize: none;
  border-radius: 5px;
  outline: none;
  padding: 1rem;
`;

export const Button = styled.button`
  display: flex;
  background-color: #00b3b3;
  padding: 0.5rem;
  text-align: center;
  align-self: center;
  border-radius: 5px;
  border: none;
  color: #f2f2f2;
  font-size: 1rem;
  font-weight: bold;
  cursor: pointer;

  &:hover {
    background-color: #006666;
  }
`;

export const StyledHeader = styled.h2`
  margin-top: 8rem;
  color: #ffffff;
`;

export const AutocompletionWrapper = styled.div`
  display: flex;
  flex-direction: row;
  gap: 0.5rem;
`;
