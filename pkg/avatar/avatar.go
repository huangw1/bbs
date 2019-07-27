/**
 * @Author: huangw1
 * @Date: 2019/7/25 15:34
 */

package avatar

// todo: replace avatars
var DefaultAvatars = []string{
	"https://file.mlog.club/avatar/club_default_avatar1.png",
	"https://file.mlog.club/avatar/club_default_avatar2.png",
	"https://file.mlog.club/avatar/club_default_avatar3.png",
	"https://file.mlog.club/avatar/club_default_avatar4.png",
	"https://file.mlog.club/avatar/club_default_avatar5.png",
	"https://file.mlog.club/avatar/club_default_avatar6.png",
}

func GetDefaultAvatar(id int64) string {
	if id == 0 {
		return DefaultAvatars[0]
	}
	index := int(id) % len(DefaultAvatars)
	return DefaultAvatars[index]
}

func IsDefaultAvatar(avatar string) bool {
	for _, defaultAvatar := range DefaultAvatars {
		if defaultAvatar == avatar {
			return true
		}
	}
	return false
}
