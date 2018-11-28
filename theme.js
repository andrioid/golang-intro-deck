import theme from "mdx-deck/themes";
import codeTheme from "prism-react-renderer/themes/vsDark";
import prismStyle from "react-syntax-highlighter/styles/prism/okaidia";

export default {
  ...theme,
  font: "Montserrat, sans-serif",
  colors: {
    ...theme.colors,
    text: "#fff",
    background: "#046065",
    link: "#ffffff",
    blockquote: "#fff",
    code: "#fff"
  },
  prism: {
    style: prismStyle
  },
  codeSurfer: {
    ...codeTheme
  }
  // Customize your presentation theme here.
  //
  // Read the docs for more info:
  // https://github.com/jxnblk/mdx-deck/blob/master/docs/theming.md
  // https://github.com/jxnblk/mdx-deck/blob/master/docs/themes.md
};
