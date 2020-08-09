package unity

import (
	"fmt"
)

func Choose(x string) string {

	if x == "h" {
		fmt.Println("有关某个命令的详细信息，请键入 HELP 命令名")
		fmt.Println("ASSOC          显示或修改文件扩展名关联。")
		fmt.Println("ATTRIB         显示或更改文件属性。")
		fmt.Println("BREAK          设置或清除扩展式 CTRL+C 检查。")
		fmt.Println("BCDEDIT        设置启动数据库中的属性以控制启动加载。")
		fmt.Println("CACLS          显示或修改文件的访问控制列表(ACL)。")
		fmt.Println("CD             显示当前目录的名称或将其更改。")
		fmt.Println("CHCP           显示或设置活动代码页数。")
		fmt.Println("CHDIR          显示当前目录的名称或将其更改。")
		fmt.Println("CHKDSK         检查磁盘并显示状态报告。")
		fmt.Println("CHKNTFS        显示或修改启动时间磁盘检查。")
		fmt.Println("CLS            清除屏幕。")
		fmt.Println("CMD            打开另一个 Windows 命令解释程序窗口。")
		fmt.Println("COLOR          设置默认控制台前景和背景颜色。")
		fmt.Println("COMP           比较两个或两套文件的内容。")
		fmt.Println("COMPACT        显示或更改 NTFS 分区上文件的压缩。")
		fmt.Println("CONVERT        将 FAT 卷转换成 NTFS。你不能转换当前驱动器。")
		fmt.Println("COPY           将至少一个文件复制到另一个位置。")
		fmt.Println("DATE           显示或设置日期。")
		fmt.Println("DEL            删除至少一个文件。")
		fmt.Println("DIR            显示一个目录中的文件和子目录。")
		fmt.Println("DISKPART       显示或配置磁盘分区属性。")
		fmt.Println("DOSKEY         编辑命令行、撤回 Windows 命令并创建宏。")
		fmt.Println("DRIVERQUERY    显示当前设备驱动程序状态和属性。")
		fmt.Println("ECHO           显示消息，或将命令回显打开或关闭。")
		fmt.Println("ENDLOCAL       结束批文件中环境更改的本地化。")
		fmt.Println("ERASE          删除一个或多个文件。")
		fmt.Println("EXIT           退出 CMD.EXE 程序(命令解释程序)。")
		fmt.Println("FC             比较两个文件或两个文件集并显示它们之间的不同。")
		fmt.Println("FIND           在一个或多个文件中搜索一个文本字符串。")
		fmt.Println("FINDSTR        在多个文件中搜索字符串。")
		fmt.Println("FOR            为一组文件中的每个文件运行一个指定的命令。")
		fmt.Println("FORMAT         格式化磁盘，以便用于 Windows。")
		fmt.Println("FSUTIL         显示或配置文件系统属性。")
		fmt.Println("FTYPE          显示或修改在文件扩展名关联中使用的文件类型。")
		fmt.Println("GOTO           将 Windows 命令解释程序定向到批处理程序中某个带标签的行。")
		fmt.Println("GPRESULT       显示计算机或用户的组策略信息。")
		fmt.Println("GRAFTABL       使 Windows 在图形模式下显示扩展字符集。")
		fmt.Println("HELP           提供 Windows 命令的帮助信息。")
		fmt.Println("IF             在批处理程序中执行有条件的处理操作。")
		fmt.Println("LABEL          创建、更改或删除磁盘的卷标。")
		fmt.Println("MD             创建一个目录。")
		fmt.Println("MKDIR          创建一个目录。")
		fmt.Println("MKLINK         创建符号链接和硬链接")
		fmt.Println("MODE           配置系统设备。")
		fmt.Println("MORE           逐屏显示输出。")
		fmt.Println("MOVE           将一个或多个文件从一个目录移动到另一个目录。")
		fmt.Println("OPENFILES      显示远程用户为了文件共享而打开的文件。")
		fmt.Println("PATH           为可执行文件显示或设置搜索路径。")
		fmt.Println("PAUSE          暂停批处理文件的处理并显示消息。")
		fmt.Println("POPD           还原通过 PUSHD 保存的当前目录的上一值。")
		fmt.Println("PRINT          打印一个文本文件。")
		fmt.Println("PROMPT         更改 Windows 命令提示。")
		fmt.Println("PUSHD          保存当前目录，然后对其进行更改。")
		fmt.Println("RD             删除目录。")
		fmt.Println("RECOVER        从损坏的或有缺陷的磁盘中恢复可读信息。")
		fmt.Println("REM            记录批处理文件或 CONFIG.SYS 中的注释(批注)。")
		fmt.Println("REN            重命名文件。")
		fmt.Println("RENAME         重命名文件。")
		fmt.Println("REPLACE        替换文件。")
		fmt.Println("RMDIR          删除目录。")
		fmt.Println("ROBOCOPY       复制文件和目录树的高级实用工具")
		fmt.Println("SET            显示、设置或删除 Windows 环境变量。")
		fmt.Println("SETLOCAL       开始本地化批处理文件中的环境更改。")
		fmt.Println("SC             显示或配置服务(后台进程)。")
		fmt.Println("SCHTASKS       安排在一台计算机上运行命令和程序。")
		fmt.Println("SHIFT          调整批处理文件中可替换参数的位置。")
		fmt.Println("SHUTDOWN       允许通过本地或远程方式正确关闭计算机。")
		fmt.Println("SORT           对输入排序。")
		fmt.Println("START          启动单独的窗口以运行指定的程序或命令。")
		fmt.Println("SUBST          将路径与驱动器号关联。")
		fmt.Println("SYSTEMINFO     显示计算机的特定属性和配置。")
		fmt.Println("TASKLIST       显示包括服务在内的所有当前运行的任务。")
		fmt.Println("TASKKILL       中止或停止正在运行的进程或应用程序。")
		fmt.Println("TIME           显示或设置系统时间。")
		fmt.Println("TITLE          设置 CMD.EXE 会话的窗口标题。")
		fmt.Println("TREE           以图形方式显示驱动程序或路径的目录结构。")
		fmt.Println("TYPE           显示文本文件的内容。")
		fmt.Println("VER            显示 Windows 的版本。")
		fmt.Println("VERIFY         告诉 Windows 是否进行验证，以确保文件正确写入磁盘。")
		fmt.Println("VOL            显示磁盘卷标和序列号。")
		fmt.Println("XCOPY          复制文件和目录树。")
		fmt.Println("WMIC           在交互式命令 shell 中显示 WMI 信息。")

		fmt.Println("有关工具的详细信息，请参阅联机帮助中的命令行参考。")
		return ""
	} else {
		fmt.Println("")
		return ""
	}
}
