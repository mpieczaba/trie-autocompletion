import React, { Component, ChangeEvent, useState } from "react";

import { TrieNode } from "../../lib";

import {
  AutocompletionWrapper,
  Button,
  InputWrapper,
  StyledHeader,
  StyledInput,
} from "./styles";
import Autocompletion from "../Autocompletion";

export class Input extends Component<
  {},
  { word: string; autocompletion: string[] }
> {
  private root = new TrieNode();

  state = {
    word: "",
    autocompletion: new Array(),
  };

  handleInput = (e: ChangeEvent<HTMLInputElement>) => {
    const word = e.target.value.split(" ").at(-1);

    this.setState({
      word: word || "",
      autocompletion: word ? this.root.search(word) : [],
    });
  };

  handleClick = () => {
    this.root.insert(this.state.word);

    this.setState({
      word: "",
      autocompletion: this.root.search(this.state.word),
    });
  };

  render(): React.ReactNode {
    return (
      <>
        <StyledHeader>Type something cool</StyledHeader>
        <InputWrapper>
          <StyledInput onChange={this.handleInput}></StyledInput>
          <Button onClick={this.handleClick}>Add to your dictionary</Button>
        </InputWrapper>

        <AutocompletionWrapper>
          {this.state.autocompletion.map((autocompletion) => (
            <Autocompletion>{autocompletion}</Autocompletion>
          ))}
        </AutocompletionWrapper>
      </>
    );
  }
}
