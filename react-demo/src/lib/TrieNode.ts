export class TrieNode {
  private children: Map<String, TrieNode>;
  private eow: boolean;

  constructor() {
    this.children = new Map();
    this.eow = false;
  }

  /**
   * Add a word to the dictionary
   * @param key word to add
   */
  public insert(key: string) {
    let node: TrieNode | undefined = this;

    for (let ch of [...key]) {
      if (!node?.children.get(ch)) {
        node?.children.set(ch, new TrieNode());
      }

      node = node?.children.get(ch);
    }

    node!.eow = true;
  }

  /**
   * Searches for words that either start or are equal to the prefix
   * @param key prefix to search for
   * @returns an array of words that either start or are equal to the prefix
   */
  public search(key: string): string[] {
    let node: TrieNode | undefined = this;

    for (let ch of [...key]) {
      if (!node?.children.get(ch)) return [];

      node = node.children.get(ch);
    }

    return node?.eow ? [key, ...node!.traverse(key)] : node!.traverse(key);
  }

  /**
   * Traverses through the dictionary from the given node and return all its words
   * @param prefix prefix of the words stored as a child nodes
   * @returns an array of words within the dictionary
   */
  private traverse(prefix: string): string[] {
    let results = new Array<string>();

    this.children.forEach((node, key) => {
      if (node.eow) results.push(prefix + key);

      results.push(...node.traverse(prefix + key));
    });

    return results;
  }
}
