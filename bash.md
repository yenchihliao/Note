* `ctrl+r`: search in history.
* `!<element>`: search the latest \<element> that appears in the first place of history and repeat the whole command.
* `!?<element>?`: search the latest \<element> that appears in the history and repeat the whole command.
* `!<count>`: use the \<count>-th command in history(negative number supported).
* `!<count>:<designator>:<modifier>`:
  * \<designator>: 
    * <number>: argv of the \<count>-th history command, which starts from *0*.
    * `^`, `$`, `*`: for the first, the last, and all the argv aside from the first respectively.
    * <range>: previous designators connected by `-`.
  * \<modifier>:
    * `h`: head
    * `t`: tail
    * `p`: print 
    * [more](https://www.gnu.org/software/bash/manual/html_node/Modifiers.html)

* ref: [How To Use Bash History Commands and Expansions on a Linux VPS](https://www.digitalocean.com/community/tutorials/how-to-use-bash-history-commands-and-expansions-on-a-linux-vps)
