# Removes Empty Folders
> It is preferable to use `find` when performing this task:
> e.g.
>
> `find . -type d -empty -delete`
>
> `find . -depth -type d -size -100k -exec rm -rv {} \`



## Usage
Default behaivor:
- lists empty folders
- y/N prompt to remove listed folders

### Flags
| Flag    | Type   | Desc.                                           |
| --------- | ------ | ---------------------------------------------- |
| Threshold | `int`  | Delete folders lower than a threshold in bytes |
| NoConfirm | `bool` | Do not ask for deletion confirmation           |


### Examples
Remove folders with size lower than 2 KiB
`rmem -T 2048`
