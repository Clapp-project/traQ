package v1

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"image"
	"io"
	"math"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/disintegration/imaging"
	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"
	"github.com/traPtitech/traQ/model"
	"github.com/traPtitech/traQ/rbac/role"
	"github.com/traPtitech/traQ/repository"
	"github.com/traPtitech/traQ/utils"
	"github.com/traPtitech/traQ/utils/ioext"
	"github.com/traPtitech/traQ/utils/set"
	"github.com/traPtitech/traQ/utils/storage"
	"github.com/traPtitech/traQ/utils/validator"
	"golang.org/x/sync/errgroup"
	"gopkg.in/guregu/null.v3"
)

var (
	dmChannelRootUUID  = uuid.Must(uuid.FromString(model.DirectMessageChannelRootID))
	pubChannelRootUUID = uuid.Nil
)

type TestRepository struct {
	FS                        storage.FileStorage
	Users                     map[uuid.UUID]model.User
	UsersLock                 sync.RWMutex
	UserGroups                map[uuid.UUID]model.UserGroup
	UserGroupsLock            sync.RWMutex
	UserGroupMembers          map[uuid.UUID]map[uuid.UUID]bool
	UserGroupMembersLock      sync.RWMutex
	UserGroupAdmins           map[uuid.UUID]map[uuid.UUID]bool
	UserGroupAdminsLock       sync.RWMutex
	Tags                      map[uuid.UUID]model.Tag
	TagsLock                  sync.RWMutex
	UserTags                  map[uuid.UUID]map[uuid.UUID]model.UsersTag
	UserTagsLock              sync.RWMutex
	Channels                  map[uuid.UUID]model.Channel
	ChannelsLock              sync.RWMutex
	ChannelSubscribes         map[uuid.UUID]map[uuid.UUID]model.ChannelSubscribeLevel
	ChannelSubscribesLock     sync.RWMutex
	PrivateChannelMembers     map[uuid.UUID]map[uuid.UUID]bool
	PrivateChannelMembersLock sync.RWMutex
	Messages                  map[uuid.UUID]model.Message
	MessagesLock              sync.RWMutex
	MessageUnreads            map[uuid.UUID]map[uuid.UUID]bool
	MessageUnreadsLock        sync.RWMutex
	MessageReports            []model.MessageReport
	MessageReportsLock        sync.RWMutex
	Pins                      map[uuid.UUID]model.Pin
	PinsLock                  sync.RWMutex
	Stars                     map[uuid.UUID]map[uuid.UUID]bool
	StarsLock                 sync.RWMutex
	Stamps                    map[uuid.UUID]model.Stamp
	StampsLock                sync.RWMutex
	Files                     map[uuid.UUID]model.File
	FilesLock                 sync.RWMutex
	FilesACL                  map[uuid.UUID]map[uuid.UUID]bool
	FilesACLLock              sync.RWMutex
	Webhooks                  map[uuid.UUID]model.WebhookBot
	WebhooksLock              sync.RWMutex
	OAuth2Clients             map[string]model.OAuth2Client
	OAuth2ClientsLock         sync.RWMutex
	OAuth2Authorizes          map[string]model.OAuth2Authorize
	OAuth2AuthorizesLock      sync.RWMutex
	OAuth2Tokens              map[uuid.UUID]model.OAuth2Token
	OAuth2TokensLock          sync.RWMutex
}

func (repo *TestRepository) LinkExternalUserAccount(userID uuid.UUID, args repository.LinkExternalUserAccountArgs) error {
	panic("implement me")
}

func (repo *TestRepository) GetLinkedExternalUserAccounts(userID uuid.UUID) ([]*model.ExternalProviderUser, error) {
	panic("implement me")
}

func (repo *TestRepository) UnlinkExternalUserAccount(userID uuid.UUID, providerName string) error {
	panic("implement me")
}

func (repo *TestRepository) GetChannelStats(channelID uuid.UUID) (*repository.ChannelStats, error) {
	panic("implement me")
}

func (repo *TestRepository) Channel(path string) (uuid.UUID, bool) {
	levels := strings.Split(path, "/")
	if len(levels) == 0 {
		return uuid.Nil, false
	}

	repo.ChannelsLock.RLock()
	defer repo.ChannelsLock.RUnlock()
	var c model.Channel
	for _, name := range levels {
		ok := false
		for _, v := range repo.Channels {
			if v.ParentID == c.ID && strings.ToLower(v.Name) == name {
				c = v
				ok = true
				break
			}
		}
		if !ok {
			return uuid.Nil, false
		}
	}
	return c.ID, true
}

func (repo *TestRepository) Group(name string) (uuid.UUID, bool) {
	repo.UserGroupsLock.RLock()
	defer repo.UserGroupsLock.RUnlock()
	for _, v := range repo.UserGroups {
		if v.Name == name {
			return v.ID, true
		}
	}
	return uuid.Nil, false
}

func (repo *TestRepository) User(name string) (uuid.UUID, bool) {
	repo.UsersLock.RLock()
	defer repo.UsersLock.RUnlock()
	for _, u := range repo.Users {
		if strings.ToLower(u.Name) == name {
			return u.ID, true
		}
	}
	return uuid.Nil, false
}

func (repo *TestRepository) GetChannelEvents(query repository.ChannelEventsQuery) (events []*model.ChannelEvent, more bool, err error) {
	panic("implement me")
}

func (repo *TestRepository) GetRole(role string) (*model.UserRole, error) {
	panic("implement me")
}

func (repo *TestRepository) CreateRole(name string) error {
	panic("implement me")
}

func (repo *TestRepository) UpdateRole(role string, args repository.UpdateRoleArgs) error {
	panic("implement me")
}

func (repo *TestRepository) GetAllRoles() ([]*model.UserRole, error) {
	return role.SystemRoles(), nil
}

func (repo *TestRepository) GetUserUnreadChannels(userID uuid.UUID) ([]*repository.UserUnreadChannel, error) {
	panic("implement me")
}

func (repo *TestRepository) GetBotByBotUserID(id uuid.UUID) (*model.Bot, error) {
	panic("implement me")
}

func (repo *TestRepository) UpdateBot(id uuid.UUID, args repository.UpdateBotArgs) error {
	panic("implement me")
}

func (repo *TestRepository) GetBotEventLogs(botID uuid.UUID, limit, offset int) ([]*model.BotEventLog, error) {
	panic("implement me")
}

func (repo *TestRepository) WriteBotEventLog(log *model.BotEventLog) error {
	panic("implement me")
}

func (repo *TestRepository) ReissueBotTokens(id uuid.UUID) (*model.Bot, error) {
	panic("implement me")
}

func (repo *TestRepository) CreateStampPalette(name, description string, stamps model.UUIDs, userID uuid.UUID) (*model.StampPalette, error) {
	panic("implement me")
}

func (repo *TestRepository) UpdateStampPalette(id uuid.UUID, args repository.UpdateStampPaletteArgs) error {
	panic("implement me")
}

func (repo *TestRepository) GetStampPalette(id uuid.UUID) (sp *model.StampPalette, err error) {
	panic("implement me")
}

func (repo *TestRepository) DeleteStampPalette(id uuid.UUID) (err error) {
	panic("implement me")
}

func (repo *TestRepository) GetStampPalettes(userID uuid.UUID) (sps []*model.StampPalette, err error) {
	panic("implement me")
}

func (repo *TestRepository) ExistStamps(stampIDs []uuid.UUID) (err error) {
	panic("implement me")
}

func NewTestRepository() *TestRepository {
	r := &TestRepository{
		FS:                    storage.NewInMemoryFileStorage(),
		Users:                 map[uuid.UUID]model.User{},
		UserGroups:            map[uuid.UUID]model.UserGroup{},
		UserGroupMembers:      map[uuid.UUID]map[uuid.UUID]bool{},
		UserGroupAdmins:       map[uuid.UUID]map[uuid.UUID]bool{},
		Tags:                  map[uuid.UUID]model.Tag{},
		UserTags:              map[uuid.UUID]map[uuid.UUID]model.UsersTag{},
		Channels:              map[uuid.UUID]model.Channel{},
		ChannelSubscribes:     map[uuid.UUID]map[uuid.UUID]model.ChannelSubscribeLevel{},
		PrivateChannelMembers: map[uuid.UUID]map[uuid.UUID]bool{},
		Messages:              map[uuid.UUID]model.Message{},
		MessageUnreads:        map[uuid.UUID]map[uuid.UUID]bool{},
		MessageReports:        []model.MessageReport{},
		Pins:                  map[uuid.UUID]model.Pin{},
		Stars:                 map[uuid.UUID]map[uuid.UUID]bool{},
		Stamps:                map[uuid.UUID]model.Stamp{},
		Files:                 map[uuid.UUID]model.File{},
		FilesACL:              map[uuid.UUID]map[uuid.UUID]bool{},
		Webhooks:              map[uuid.UUID]model.WebhookBot{},
		OAuth2Clients:         map[string]model.OAuth2Client{},
		OAuth2Authorizes:      map[string]model.OAuth2Authorize{},
		OAuth2Tokens:          map[uuid.UUID]model.OAuth2Token{},
	}
	_, _ = r.CreateUser(repository.CreateUserArgs{Name: "traq", Password: "traq", Role: role.Admin})
	return r
}

func (repo *TestRepository) Sync() (bool, error) {
	panic("implement me")
}

func (repo *TestRepository) GetFS() storage.FileStorage {
	return repo.FS
}

func (repo *TestRepository) CreateUser(args repository.CreateUserArgs) (model.UserInfo, error) {
	repo.UsersLock.Lock()
	defer repo.UsersLock.Unlock()

	uid := uuid.Must(uuid.NewV4())
	user := model.User{
		ID:          uid,
		Name:        args.Name,
		DisplayName: args.DisplayName,
		Status:      model.UserAccountStatusActive,
		Bot:         false,
		Role:        args.Role,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Profile: &model.UserProfile{
			UserID:    uid,
			UpdatedAt: time.Now(),
		},
	}

	if len(args.Password) > 0 {
		salt := utils.GenerateSalt()
		user.Password = hex.EncodeToString(utils.HashPassword(args.Password, salt))
		user.Salt = hex.EncodeToString(salt)
	}

	if args.IconFileID.Valid {
		user.Icon = args.IconFileID.UUID
	} else {
		iconID, err := repository.GenerateIconFile(repo, user.Name)
		if err != nil {
			return nil, err
		}
		user.Icon = iconID
	}

	if args.ExternalLogin != nil {
		panic("implement me")
	}

	for _, v := range repo.Users {
		if v.Name == user.Name {
			return nil, repository.ErrAlreadyExists
		}
	}

	repo.Users[user.ID] = user
	return &user, nil
}

func (repo *TestRepository) GetUser(id uuid.UUID, withProfile bool) (model.UserInfo, error) {
	repo.UsersLock.RLock()
	u, ok := repo.Users[id]
	repo.UsersLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &u, nil
}

func (repo *TestRepository) GetUserByName(name string, withProfile bool) (model.UserInfo, error) {
	repo.UsersLock.RLock()
	defer repo.UsersLock.RUnlock()
	for _, u := range repo.Users {
		u := u
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (repo *TestRepository) GetUserByExternalID(providerName, externalID string, withProfile bool) (model.UserInfo, error) {
	panic("implement me")
}

func (repo *TestRepository) GetUsers(query repository.UsersQuery) ([]model.UserInfo, error) {
	result := make([]model.UserInfo, 0, len(repo.Users))
	repo.UsersLock.RLock()
	repo.PrivateChannelMembersLock.RLock()
	repo.UserGroupMembersLock.RLock()
	for _, u := range repo.Users {
		if query.IsBot.Valid {
			if u.Bot != query.IsBot.Bool {
				continue
			}
		}
		if query.IsActive.Valid {
			if query.IsActive.Bool {
				if u.Status != model.UserAccountStatusActive {
					continue
				}
			} else {
				if u.Status == model.UserAccountStatusActive {
					continue
				}
			}
		}
		if query.IsCMemberOf.Valid {
			arr, ok := repo.PrivateChannelMembers[query.IsCMemberOf.UUID]
			if !ok || !arr[u.ID] {
				continue
			}
		}
		if query.IsGMemberOf.Valid {
			arr, ok := repo.UserGroupMembers[query.IsGMemberOf.UUID]
			if !ok || !arr[u.ID] {
				continue
			}
		}
		u := u
		result = append(result, &u)
	}
	repo.UserGroupMembersLock.RUnlock()
	repo.PrivateChannelMembersLock.RUnlock()
	repo.UsersLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) GetUserIDs(query repository.UsersQuery) ([]uuid.UUID, error) {
	ids := make([]uuid.UUID, 0)
	repo.UsersLock.RLock()
	repo.PrivateChannelMembersLock.RLock()
	repo.UserGroupMembersLock.RLock()
	for _, v := range repo.Users {
		if query.IsBot.Valid {
			if v.Bot != query.IsBot.Bool {
				continue
			}
		}
		if query.IsActive.Valid {
			if query.IsActive.Bool {
				if v.Status != model.UserAccountStatusActive {
					continue
				}
			} else {
				if v.Status == model.UserAccountStatusActive {
					continue
				}
			}
		}
		if query.IsCMemberOf.Valid {
			arr, ok := repo.PrivateChannelMembers[query.IsCMemberOf.UUID]
			if !ok || !arr[v.ID] {
				continue
			}
		}
		if query.IsGMemberOf.Valid {
			arr, ok := repo.UserGroupMembers[query.IsGMemberOf.UUID]
			if !ok || !arr[v.ID] {
				continue
			}
		}
		ids = append(ids, v.ID)
	}
	repo.UserGroupMembersLock.RUnlock()
	repo.PrivateChannelMembersLock.RUnlock()
	repo.UsersLock.RUnlock()
	return ids, nil
}

func (repo *TestRepository) UserExists(id uuid.UUID) (bool, error) {
	repo.UsersLock.RLock()
	_, ok := repo.Users[id]
	repo.UsersLock.RUnlock()
	return ok, nil
}

func (repo *TestRepository) UpdateUser(id uuid.UUID, args repository.UpdateUserArgs) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UsersLock.Lock()
	defer repo.UsersLock.Unlock()

	u, ok := repo.Users[id]
	if !ok {
		return repository.ErrNotFound
	}

	if args.DisplayName.Valid {
		if utf8.RuneCountInString(args.DisplayName.String) > 64 {
			return repository.ArgError("args.DisplayName", "DisplayName must be shorter than 64 characters")
		}
		u.DisplayName = args.DisplayName.String
		u.UpdatedAt = time.Now()
	}
	if args.Password.Valid {
		salt := utils.GenerateSalt()
		hashed := utils.HashPassword(args.Password.String, salt)
		u.Salt = hex.EncodeToString(salt)
		u.Password = hex.EncodeToString(hashed)
		u.UpdatedAt = time.Now()
	}
	if args.TwitterID.Valid {
		if len(args.TwitterID.String) > 0 && !validator.TwitterIDRegex.MatchString(args.TwitterID.String) {
			return repository.ArgError("args.TwitterID", "invalid TwitterID")
		}
		u.Profile.TwitterID = args.TwitterID.String
		u.Profile.UpdatedAt = time.Now()
	}
	if args.Role.Valid {
		u.Role = args.Role.String
		u.UpdatedAt = time.Now()
	}
	if args.IconFileID.Valid {
		u.Icon = args.IconFileID.UUID
		u.UpdatedAt = time.Now()
	}
	if args.LastOnline.Valid {
		u.Profile.LastOnline = args.LastOnline
		u.Profile.UpdatedAt = time.Now()
	}

	repo.Users[id] = u
	return nil
}

func (repo *TestRepository) CreateUserGroup(name, description, gType string, adminID uuid.UUID) (*model.UserGroup, error) {
	g := model.UserGroup{
		ID:          uuid.Must(uuid.NewV4()),
		Name:        name,
		Description: description,
		Type:        gType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repo.UserGroupsLock.Lock()
	repo.UserGroupAdminsLock.Lock()
	defer repo.UserGroupsLock.Unlock()
	defer repo.UserGroupAdminsLock.Unlock()

	// 名前チェック
	if len(g.Name) == 0 || utf8.RuneCountInString(g.Name) > 30 {
		return nil, repository.ArgError("name", "Name must be non-empty and shorter than 31 characters")
	}
	// タイプチェック
	if utf8.RuneCountInString(g.Type) > 30 {
		return nil, repository.ArgError("Type", "Type must be shorter than 31 characters")
	}

	for _, v := range repo.UserGroups {
		if v.Name == name {
			return nil, repository.ErrAlreadyExists
		}
	}
	repo.UserGroups[g.ID] = g
	repo.UserGroupAdmins[g.ID] = make(map[uuid.UUID]bool)
	repo.UserGroupAdmins[g.ID][adminID] = true
	g.Members = make([]*model.UserGroupMember, 0)
	g.Admins = []*model.UserGroupAdmin{{GroupID: g.ID, UserID: adminID}}
	return &g, nil
}

func (repo *TestRepository) UpdateUserGroup(id uuid.UUID, args repository.UpdateUserGroupNameArgs) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}

	repo.UserGroupsLock.Lock()
	repo.UsersLock.RLock()
	defer repo.UserGroupsLock.Unlock()
	defer repo.UsersLock.RUnlock()
	g, ok := repo.UserGroups[id]
	if !ok {
		return repository.ErrNotFound
	}
	changed := false
	if args.Name.Valid {
		if len(args.Name.String) == 0 || utf8.RuneCountInString(args.Name.String) > 30 {
			return repository.ArgError("args.Name", "Name must be non-empty and shorter than 31 characters")
		}

		for _, v := range repo.UserGroups {
			if v.Name == args.Name.String {
				return repository.ErrAlreadyExists
			}
		}
		g.Name = args.Name.String
	}
	if args.Description.Valid {
		g.Description = args.Description.String
		changed = true
	}
	if args.Type.Valid {
		if utf8.RuneCountInString(args.Type.String) > 30 {
			return repository.ArgError("args.Type", "Type must be shorter than 31 characters")
		}
		g.Type = args.Type.String
		changed = true
	}

	if changed {
		g.UpdatedAt = time.Now()
		repo.UserGroups[id] = g
	}
	return nil
}

func (repo *TestRepository) DeleteUserGroup(id uuid.UUID) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserGroupsLock.Lock()
	defer repo.UserGroupsLock.Unlock()
	repo.UserGroupMembersLock.Lock()
	defer repo.UserGroupMembersLock.Unlock()
	repo.UserGroupAdminsLock.Lock()
	defer repo.UserGroupAdminsLock.Unlock()
	if _, ok := repo.UserGroups[id]; !ok {
		return repository.ErrNotFound
	}
	delete(repo.UserGroups, id)
	delete(repo.UserGroupMembers, id)
	delete(repo.UserGroupAdmins, id)
	return nil
}

func (repo *TestRepository) GetUserGroup(id uuid.UUID) (*model.UserGroup, error) {
	if id == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	repo.UserGroupsLock.RLock()
	repo.UserGroupAdminsLock.Lock()
	repo.UserGroupMembersLock.Lock()
	defer repo.UserGroupsLock.RUnlock()
	defer repo.UserGroupAdminsLock.Unlock()
	defer repo.UserGroupMembersLock.Unlock()
	g, ok := repo.UserGroups[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	members := repo.UserGroupMembers[id]
	for u := range members {
		g.Members = append(g.Members, &model.UserGroupMember{GroupID: g.ID, UserID: u})
	}
	admins := repo.UserGroupAdmins[id]
	for u := range admins {
		g.Admins = append(g.Admins, &model.UserGroupAdmin{GroupID: g.ID, UserID: u})
	}
	return &g, nil
}

func (repo *TestRepository) GetUserGroupByName(name string) (*model.UserGroup, error) {
	if len(name) == 0 {
		return nil, repository.ErrNotFound
	}
	repo.UserGroupsLock.RLock()
	repo.UserGroupAdminsLock.Lock()
	repo.UserGroupMembersLock.Lock()
	defer repo.UserGroupsLock.RUnlock()
	defer repo.UserGroupAdminsLock.Unlock()
	defer repo.UserGroupMembersLock.Unlock()
	for _, v := range repo.UserGroups {
		if v.Name == name {
			members := repo.UserGroupMembers[v.ID]
			for u := range members {
				v.Members = append(v.Members, &model.UserGroupMember{GroupID: v.ID, UserID: u})
			}
			admins := repo.UserGroupAdmins[v.ID]
			for u := range admins {
				v.Admins = append(v.Admins, &model.UserGroupAdmin{GroupID: v.ID, UserID: u})
			}
			return &v, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (repo *TestRepository) GetUserBelongingGroupIDs(userID uuid.UUID) ([]uuid.UUID, error) {
	groups := make([]uuid.UUID, 0)
	if userID == uuid.Nil {
		return groups, nil
	}
	repo.UserGroupMembersLock.RLock()
	for gid, users := range repo.UserGroupMembers {
		for uid := range users {
			if uid == userID {
				groups = append(groups, gid)
				break
			}
		}
	}
	repo.UserGroupMembersLock.RUnlock()
	return groups, nil
}

func (repo *TestRepository) GetAllUserGroups() ([]*model.UserGroup, error) {
	groups := make([]*model.UserGroup, 0)
	repo.UserGroupsLock.RLock()
	repo.UserGroupAdminsLock.Lock()
	repo.UserGroupMembersLock.Lock()
	for _, v := range repo.UserGroups {
		v := v
		members := repo.UserGroupMembers[v.ID]
		for u := range members {
			v.Members = append(v.Members, &model.UserGroupMember{GroupID: v.ID, UserID: u})
		}
		admins := repo.UserGroupAdmins[v.ID]
		for u := range admins {
			v.Admins = append(v.Admins, &model.UserGroupAdmin{GroupID: v.ID, UserID: u})
		}
		groups = append(groups, &v)
	}
	repo.UserGroupMembersLock.Unlock()
	repo.UserGroupAdminsLock.Unlock()
	repo.UserGroupsLock.RUnlock()
	return groups, nil
}

func (repo *TestRepository) AddUserToGroup(userID, groupID uuid.UUID, role string) error {
	if userID == uuid.Nil || groupID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserGroupsLock.Lock()
	defer repo.UserGroupsLock.Unlock()
	repo.UserGroupMembersLock.Lock()
	defer repo.UserGroupMembersLock.Unlock()
	g, ok := repo.UserGroups[groupID]
	if !ok {
		return nil
	}
	users, ok := repo.UserGroupMembers[groupID]
	if !ok {
		users = make(map[uuid.UUID]bool)
		repo.UserGroupMembers[groupID] = users
	}
	if !users[userID] {
		users[userID] = true
		g.UpdatedAt = time.Now()
		repo.UserGroups[groupID] = g
	}
	return nil
}

func (repo *TestRepository) RemoveUserFromGroup(userID, groupID uuid.UUID) error {
	if userID == uuid.Nil || groupID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserGroupsLock.Lock()
	defer repo.UserGroupsLock.Unlock()
	repo.UserGroupMembersLock.Lock()
	defer repo.UserGroupMembersLock.Unlock()
	g, ok := repo.UserGroups[groupID]
	if !ok {
		return nil
	}

	users, ok := repo.UserGroupMembers[groupID]
	if ok && users[userID] {
		delete(users, userID)
		g.UpdatedAt = time.Now()
		repo.UserGroups[groupID] = g
	}
	return nil
}

func (repo *TestRepository) AddUserToGroupAdmin(userID, groupID uuid.UUID) error {
	if userID == uuid.Nil || groupID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserGroupsLock.Lock()
	defer repo.UserGroupsLock.Unlock()
	repo.UserGroupAdminsLock.Lock()
	defer repo.UserGroupAdminsLock.Unlock()
	g, ok := repo.UserGroups[groupID]
	if !ok {
		return nil
	}
	users := repo.UserGroupAdmins[groupID]
	if !users[userID] {
		users[userID] = true
		g.UpdatedAt = time.Now()
		repo.UserGroups[groupID] = g
	}
	return nil
}

func (repo *TestRepository) RemoveUserFromGroupAdmin(userID, groupID uuid.UUID) error {
	if userID == uuid.Nil || groupID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserGroupsLock.Lock()
	defer repo.UserGroupsLock.Unlock()
	repo.UserGroupAdminsLock.Lock()
	defer repo.UserGroupAdminsLock.Unlock()
	g, ok := repo.UserGroups[groupID]
	if !ok {
		return nil
	}

	users, ok := repo.UserGroupAdmins[groupID]
	if ok && users[userID] {
		delete(users, userID)
		g.UpdatedAt = time.Now()
		repo.UserGroups[groupID] = g
	}
	return nil
}

func (repo *TestRepository) CreateTag(name string) (*model.Tag, error) {
	repo.TagsLock.Lock()
	defer repo.TagsLock.Unlock()
	// 名前チェック
	if len(name) == 0 || utf8.RuneCountInString(name) > 30 {
		return nil, repository.ArgError("name", "Name must be non-empty and shorter than 31 characters")
	}
	for _, t := range repo.Tags {
		if t.Name == name {
			return nil, repository.ErrAlreadyExists
		}
	}
	t := model.Tag{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.Tags[t.ID] = t
	return &t, nil
}

func (repo *TestRepository) GetTagByID(id uuid.UUID) (*model.Tag, error) {
	repo.TagsLock.RLock()
	t, ok := repo.Tags[id]
	repo.TagsLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &t, nil
}

func (repo *TestRepository) GetTagByName(name string) (*model.Tag, error) {
	repo.TagsLock.RLock()
	defer repo.TagsLock.RUnlock()
	for _, t := range repo.Tags {
		if t.Name == name {
			return &t, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (repo *TestRepository) GetOrCreateTagByName(name string) (*model.Tag, error) {
	if len(name) == 0 {
		return nil, repository.ErrNotFound
	}
	if utf8.RuneCountInString(name) > 30 {
		return nil, repository.ArgError("name", "tag must be non-empty and shorter than 31 characters")
	}
	repo.TagsLock.Lock()
	defer repo.TagsLock.Unlock()
	for _, t := range repo.Tags {
		if t.Name == name {
			return &t, nil
		}
	}
	t := model.Tag{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.Tags[t.ID] = t
	return &t, nil
}

func (repo *TestRepository) AddUserTag(userID, tagID uuid.UUID) error {
	if userID == uuid.Nil || tagID == uuid.Nil {
		return repository.ErrNilID
	}
	ut := model.UsersTag{
		UserID:    userID,
		TagID:     tagID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.UserTagsLock.Lock()
	tags, ok := repo.UserTags[userID]
	if !ok {
		tags = make(map[uuid.UUID]model.UsersTag)
		repo.UserTags[userID] = tags
	}
	if _, ok := tags[tagID]; ok {
		return repository.ErrAlreadyExists
	}
	tags[tagID] = ut
	repo.UserTagsLock.Unlock()
	return nil
}

func (repo *TestRepository) ChangeUserTagLock(userID, tagID uuid.UUID, locked bool) error {
	if userID == uuid.Nil || tagID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserTagsLock.Lock()
	defer repo.UserTagsLock.Unlock()
	for id, tag := range repo.UserTags[userID] {
		if id == tagID {
			tag.IsLocked = locked
			tag.UpdatedAt = time.Now()
			repo.UserTags[userID][tagID] = tag
			return nil
		}
	}
	return nil
}

func (repo *TestRepository) DeleteUserTag(userID, tagID uuid.UUID) error {
	if userID == uuid.Nil || tagID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.UserTagsLock.Lock()
	tags, ok := repo.UserTags[userID]
	if ok {
		delete(tags, tagID)
	}
	repo.UserTagsLock.Unlock()
	return nil
}

func (repo *TestRepository) GetUserTag(userID, tagID uuid.UUID) (*model.UsersTag, error) {
	repo.UserTagsLock.RLock()
	defer repo.UserTagsLock.RUnlock()
	tags, ok := repo.UserTags[userID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	ut, ok := tags[tagID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	repo.TagsLock.RLock()
	ut.Tag = repo.Tags[ut.TagID]
	repo.TagsLock.RUnlock()
	return &ut, nil
}

func (repo *TestRepository) GetUserTagsByUserID(userID uuid.UUID) ([]*model.UsersTag, error) {
	tags := make([]*model.UsersTag, 0)
	repo.UserTagsLock.RLock()
	for tid, ut := range repo.UserTags[userID] {
		ut := ut
		repo.TagsLock.RLock()
		ut.Tag = repo.Tags[tid]
		repo.TagsLock.RUnlock()
		tags = append(tags, &ut)
	}
	repo.UserTagsLock.RUnlock()
	return tags, nil
}

func (repo *TestRepository) GetUserIDsByTag(tag string) ([]uuid.UUID, error) {
	users := make([]uuid.UUID, 0)
	repo.TagsLock.RLock()
	tid := uuid.Nil
	for _, t := range repo.Tags {
		if t.Name == tag {
			tid = t.ID
		}
	}
	repo.TagsLock.RUnlock()
	if tid == uuid.Nil {
		return users, nil
	}
	repo.UserTagsLock.RLock()
	for uid, tags := range repo.UserTags {
		if _, ok := tags[tid]; ok {
			users = append(users, uid)
		}
	}
	repo.UserTagsLock.RUnlock()
	return users, nil
}

func (repo *TestRepository) GetUserIDsByTagID(tagID uuid.UUID) ([]uuid.UUID, error) {
	users := make([]uuid.UUID, 0)
	repo.UserTagsLock.RLock()
	for uid, tags := range repo.UserTags {
		if _, ok := tags[tagID]; ok {
			users = append(users, uid)
		}
	}
	repo.UserTagsLock.RUnlock()
	return users, nil
}

func (repo *TestRepository) CreatePublicChannel(name string, parent, creatorID uuid.UUID) (*model.Channel, error) {
	// チャンネル名検証
	if !validator.ChannelRegex.MatchString(name) {
		return nil, repository.ArgError("name", "invalid name")
	}
	if has, err := repo.IsChannelPresent(name, parent); err != nil {
		return nil, err
	} else if has {
		return nil, repository.ErrAlreadyExists
	}

	switch parent {
	case pubChannelRootUUID: // ルート
		break
	case dmChannelRootUUID: // DMルート
		return nil, repository.ErrForbidden
	default: // ルート以外
		pCh, err := repo.GetChannel(parent)
		if err != nil {
			return nil, err
		}

		// DMチャンネルの子チャンネルには出来ない
		if pCh.IsDMChannel() {
			return nil, repository.ErrForbidden
		}

		// 親と公開状況が一致しているか
		if !pCh.IsPublic {
			return nil, repository.ErrForbidden
		}

		// 深さを検証
		for parent, depth := pCh, 2; ; { // 祖先
			if parent.ParentID == uuid.Nil {
				// ルート
				break
			}

			parent, err = repo.GetChannel(parent.ParentID)
			if err != nil {
				if err == repository.ErrNotFound {
					break
				}
				return nil, err
			}
			depth++
			if depth > model.MaxChannelDepth {
				return nil, repository.ErrChannelDepthLimitation
			}
		}
	}

	ch := model.Channel{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      name,
		ParentID:  parent,
		CreatorID: creatorID,
		UpdaterID: creatorID,
		IsPublic:  true,
		IsForced:  false,
		IsVisible: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.ChannelsLock.Lock()
	repo.Channels[ch.ID] = ch
	repo.ChannelsLock.Unlock()
	return &ch, nil
}

func (repo *TestRepository) UpdateChannel(channelID uuid.UUID, args repository.UpdateChannelArgs) error {
	if channelID == uuid.Nil {
		return repository.ErrNilID
	}

	repo.ChannelsLock.Lock()
	defer repo.ChannelsLock.Unlock()
	ch, ok := repo.Channels[channelID]
	if !ok {
		return repository.ErrNotFound
	}

	if args.Topic.Valid {
		ch.Topic = args.Topic.String
	}
	if args.Visibility.Valid {
		ch.IsVisible = args.Visibility.Bool
	}
	if args.ForcedNotification.Valid {
		ch.IsForced = args.ForcedNotification.Bool
	}
	if args.Name.Valid || args.Parent.Valid {
		// ダイレクトメッセージチャンネルの名前・親は変更不可能
		if ch.IsDMChannel() {
			return repository.ErrForbidden
		}

		// チャンネル名重複を確認
		{
			var (
				n string
				p uuid.UUID
			)

			if args.Name.Valid {
				n = args.Name.String
			} else {
				n = ch.Name
			}
			if args.Parent.Valid {
				p = args.Parent.UUID
			} else {
				p = ch.ParentID
			}

			// チャンネル名重複を確認
			for _, c := range repo.Channels {
				if c.Name == n && c.ParentID == p {
					return repository.ErrAlreadyExists
				}
			}
		}

		if args.Name.Valid {
			// チャンネル名検証
			if !validator.ChannelRegex.MatchString(args.Name.String) {
				return repository.ArgError("args.Name", "invalid name")
			}

			ch.Name = args.Name.String
		}
		if args.Parent.Valid {
			// チャンネル階層検証
			switch args.Parent.UUID {
			case uuid.Nil:
				// ルートチャンネル
				break
			case dmChannelRootUUID:
				// DMチャンネルには出来ない
				return repository.ErrForbidden
			default:
				pCh, ok := repo.Channels[args.Parent.UUID]
				if !ok {
					return repository.ArgError("args.Name", "invalid name")
				}

				// DMチャンネルの子チャンネルには出来ない
				if pCh.IsDMChannel() {
					return repository.ArgError("args.Name", "invalid name")
				}

				// 親と公開状況が一致しているか
				if ch.IsPublic != pCh.IsPublic {
					return repository.ArgError("args.Name", "invalid name")
				}

				// 深さを検証
				depth := 1 // ↑で見た親
				for {      // 祖先
					if pCh.ParentID == uuid.Nil {
						// ルート
						break
					}
					if ch.ID == pCh.ID {
						// ループ検出
						return repository.ErrChannelDepthLimitation
					}

					pCh, ok = repo.Channels[pCh.ParentID]
					if !ok {
						break
					}
					depth++
					if depth >= model.MaxChannelDepth {
						return repository.ErrChannelDepthLimitation
					}
				}
				depth += repo.getChannelDepthWithoutLock(ch.ID) // 子孫 (自分を含む)
				if depth > model.MaxChannelDepth {
					return repository.ErrChannelDepthLimitation
				}
			}

			ch.ParentID = args.Parent.UUID
		}
	}

	ch.UpdatedAt = time.Now()
	ch.UpdaterID = args.UpdaterID
	repo.Channels[channelID] = ch
	return nil
}

func (repo *TestRepository) DeleteChannel(channelID uuid.UUID) error {
	if channelID == uuid.Nil {
		return repository.ErrNilID
	}

	desc, err := repo.GetDescendantChannelIDs(channelID)
	if err != nil {
		return err
	}
	repo.ChannelsLock.Lock()
	for _, id := range append(desc, channelID) {
		delete(repo.Channels, id)
	}
	repo.ChannelsLock.Unlock()
	return nil
}

func (repo *TestRepository) GetChannel(channelID uuid.UUID) (*model.Channel, error) {
	repo.ChannelsLock.RLock()
	ch, ok := repo.Channels[channelID]
	repo.ChannelsLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &ch, nil
}

func (repo *TestRepository) GetChannelByMessageID(messageID uuid.UUID) (*model.Channel, error) {
	repo.MessagesLock.RLock()
	m, ok := repo.Messages[messageID]
	repo.MessagesLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	repo.ChannelsLock.RLock()
	ch, ok := repo.Channels[m.ChannelID]
	repo.ChannelsLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &ch, nil
}

func (repo *TestRepository) GetChannelsByUserID(userID uuid.UUID) ([]*model.Channel, error) {
	result := make([]*model.Channel, 0)
	repo.ChannelsLock.RLock()
	for _, ch := range repo.Channels {
		ch := ch
		if ch.IsPublic {
			result = append(result, &ch)
		} else if userID != uuid.Nil {
			ok, _ := repo.IsUserPrivateChannelMember(ch.ID, userID)
			if ok {
				result = append(result, &ch)
			}
		}
	}
	repo.ChannelsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) GetDirectMessageChannel(user1, user2 uuid.UUID) (*model.Channel, error) {
	panic("implement me")
}

func (repo *TestRepository) GetDirectMessageChannelMapping(userID uuid.UUID) (map[uuid.UUID]uuid.UUID, error) {
	panic("implement me")
}

func (repo *TestRepository) IsChannelPresent(name string, parent uuid.UUID) (bool, error) {
	repo.ChannelsLock.RLock()
	defer repo.ChannelsLock.RUnlock()
	for _, ch := range repo.Channels {
		if ch.Name == name && ch.ParentID == parent {
			return true, nil
		}
	}
	return false, nil
}

func (repo *TestRepository) IsChannelAccessibleToUser(userID, channelID uuid.UUID) (bool, error) {
	if userID == uuid.Nil || channelID == uuid.Nil {
		return false, nil
	}
	repo.ChannelsLock.RLock()
	ch, ok := repo.Channels[channelID]
	repo.ChannelsLock.RUnlock()
	if !ok {
		return false, nil
	}
	if ch.IsPublic {
		return true, nil
	}
	return repo.IsUserPrivateChannelMember(channelID, userID)
}

func (repo *TestRepository) GetParentChannel(channelID uuid.UUID) (*model.Channel, error) {
	repo.ChannelsLock.RLock()
	defer repo.ChannelsLock.RUnlock()
	ch, ok := repo.Channels[channelID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	if ch.ParentID == uuid.Nil {
		return nil, nil
	}
	pCh, ok := repo.Channels[ch.ParentID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &pCh, nil
}

func (repo *TestRepository) GetChildrenChannelIDs(channelID uuid.UUID) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, 0)
	repo.ChannelsLock.RLock()
	for cid, ch := range repo.Channels {
		if ch.ParentID == channelID {
			result = append(result, cid)
		}
	}
	repo.ChannelsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) GetDescendantChannelIDs(channelID uuid.UUID) ([]uuid.UUID, error) {
	var descendants []uuid.UUID
	children, err := repo.GetChildrenChannelIDs(channelID)
	if err != nil {
		return nil, err
	}
	descendants = append(descendants, children...)
	for _, v := range children {
		sub, err := repo.GetDescendantChannelIDs(v)
		if err != nil {
			return nil, err
		}
		descendants = append(descendants, sub...)
	}
	return descendants, nil
}

func (repo *TestRepository) GetAscendantChannelIDs(channelID uuid.UUID) ([]uuid.UUID, error) {
	var ascendants []uuid.UUID
	parent, err := repo.GetParentChannel(channelID)
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, nil
		}
		return nil, err
	} else if parent == nil {
		return []uuid.UUID{}, nil
	}
	ascendants = append(ascendants, parent.ID)
	sub, err := repo.GetAscendantChannelIDs(parent.ID)
	if err != nil {
		return nil, err
	}
	ascendants = append(ascendants, sub...)
	return ascendants, nil
}

func (repo *TestRepository) GetChannelPath(id uuid.UUID) (string, error) {
	panic("implement me")
}

func (repo *TestRepository) getChannelDepthWithoutLock(id uuid.UUID) int {
	children := make([]uuid.UUID, 0)
	for cid, ch := range repo.Channels {
		if ch.ParentID == id {
			children = append(children, cid)
		}
	}
	max := 0
	for _, v := range children {
		d := repo.getChannelDepthWithoutLock(v)
		if max < d {
			max = d
		}
	}
	return max + 1
}

func (repo *TestRepository) GetPrivateChannelMemberIDs(channelID uuid.UUID) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, 0)
	repo.PrivateChannelMembersLock.RLock()
	for uid := range repo.PrivateChannelMembers[channelID] {
		result = append(result, uid)
	}
	repo.PrivateChannelMembersLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) IsUserPrivateChannelMember(channelID, userID uuid.UUID) (bool, error) {
	repo.PrivateChannelMembersLock.RLock()
	defer repo.PrivateChannelMembersLock.RUnlock()
	uids, ok := repo.PrivateChannelMembers[channelID]
	if !ok {
		return false, nil
	}
	for uid := range uids {
		if userID == uid {
			return true, nil
		}
	}
	return false, nil
}

func (repo *TestRepository) ChangeChannelSubscription(channelID uuid.UUID, args repository.ChangeChannelSubscriptionArgs) error {
	if channelID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.ChannelSubscribesLock.Lock()
	current, ok := repo.ChannelSubscribes[channelID]
	if !ok {
		current = make(map[uuid.UUID]model.ChannelSubscribeLevel)
		repo.ChannelSubscribes[channelID] = current
	}

	for uid, level := range args.Subscription {
		if cl := current[uid]; cl == level {
			continue // 既に同じ設定がされているのでスキップ
		}

		switch level {
		case model.ChannelSubscribeLevelNone:
			if _, ok := current[uid]; !ok {
				continue // 既にオフ
			}

			if args.KeepOffLevel {
				if cl := current[uid]; cl == model.ChannelSubscribeLevelMark {
					continue // 未読管理のみをキープしたままにする
				}
			}

			delete(current, uid)

		case model.ChannelSubscribeLevelMark:
			repo.UsersLock.RLock()
			_, ok := repo.Users[uid]
			repo.UsersLock.RUnlock()
			if !ok {
				continue
			}

			current[uid] = model.ChannelSubscribeLevelMark

		case model.ChannelSubscribeLevelMarkAndNotify:
			repo.UsersLock.RLock()
			_, ok := repo.Users[uid]
			repo.UsersLock.RUnlock()
			if !ok {
				continue
			}

			current[uid] = model.ChannelSubscribeLevelMarkAndNotify

		}
	}

	repo.ChannelSubscribesLock.Unlock()
	return nil
}

func (repo *TestRepository) GetChannelSubscriptions(query repository.ChannelSubscriptionQuery) ([]*model.UserSubscribeChannel, error) {
	repo.ChannelSubscribesLock.Lock()
	result := make([]*model.UserSubscribeChannel, 0)

	for cid, users := range repo.ChannelSubscribes {
		if query.ChannelID.Valid && cid != query.ChannelID.UUID {
			continue
		}
		for uid, level := range users {
			if query.UserID.Valid && uid != query.UserID.UUID {
				continue
			}

			switch query.Level {
			case model.ChannelSubscribeLevelMark:
				if level != model.ChannelSubscribeLevelMark {
					continue
				}
			case model.ChannelSubscribeLevelMarkAndNotify:
				if level != model.ChannelSubscribeLevelMarkAndNotify {
					continue
				}
			default:
				if level != model.ChannelSubscribeLevelNone {
					continue
				}
			}

			result = append(result, &model.UserSubscribeChannel{
				ChannelID: cid,
				UserID:    uid,
				Mark:      level >= model.ChannelSubscribeLevelMark,
				Notify:    level >= model.ChannelSubscribeLevelMarkAndNotify,
			})
		}
	}

	repo.ChannelSubscribesLock.Unlock()
	return result, nil
}

func (repo *TestRepository) CreateMessage(userID, channelID uuid.UUID, text string) (*model.Message, error) {
	if userID == uuid.Nil || channelID == uuid.Nil {
		return nil, repository.ErrNilID
	}

	m := &model.Message{
		ID:        uuid.Must(uuid.NewV4()),
		UserID:    userID,
		ChannelID: channelID,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Stamps:    make([]model.MessageStamp, 0),
	}

	repo.MessagesLock.Lock()
	repo.Messages[m.ID] = *m
	repo.MessagesLock.Unlock()
	return m, nil
}

func (repo *TestRepository) UpdateMessage(messageID uuid.UUID, text string) error {
	if messageID == uuid.Nil {
		return repository.ErrNilID
	}

	repo.MessagesLock.Lock()
	defer repo.MessagesLock.Unlock()
	m, ok := repo.Messages[messageID]
	if !ok {
		return repository.ErrNotFound
	}
	m.Text = text
	m.UpdatedAt = time.Now()
	repo.Messages[messageID] = m
	return nil
}

func (repo *TestRepository) DeleteMessage(messageID uuid.UUID) error {
	if messageID == uuid.Nil {
		return repository.ErrNilID
	}

	repo.MessagesLock.Lock()
	defer repo.MessagesLock.Unlock()
	if _, ok := repo.Messages[messageID]; !ok {
		return repository.ErrNotFound
	}
	delete(repo.Messages, messageID)
	return nil
}

func (repo *TestRepository) GetMessageByID(messageID uuid.UUID) (*model.Message, error) {
	repo.MessagesLock.RLock()
	m, ok := repo.Messages[messageID]
	repo.MessagesLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	m.Stamps = make([]model.MessageStamp, 0)
	return &m, nil
}

func (repo *TestRepository) GetMessages(query repository.MessagesQuery) (messages []*model.Message, more bool, err error) {
	tmp := make([]*model.Message, 0)

	repo.MessagesLock.RLock()
	if query.Channel != uuid.Nil {
		if query.User != uuid.Nil {
			for _, v := range repo.Messages {
				if v.ChannelID == query.Channel && v.UserID == query.User {
					v := v
					v.Stamps = make([]model.MessageStamp, 0)
					tmp = append(tmp, &v)
				}
			}
		} else {
			for _, v := range repo.Messages {
				if v.ChannelID == query.Channel {
					v := v
					v.Stamps = make([]model.MessageStamp, 0)
					tmp = append(tmp, &v)
				}
			}
		}
	} else if query.User != uuid.Nil {
		for _, v := range repo.Messages {
			if v.UserID == query.User {
				v := v
				v.Stamps = make([]model.MessageStamp, 0)
				tmp = append(tmp, &v)
			}
		}
	} else {
		for _, v := range repo.Messages {
			v := v
			v.Stamps = make([]model.MessageStamp, 0)
			tmp = append(tmp, &v)
		}
	}
	repo.MessagesLock.RUnlock()

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].CreatedAt.After(tmp[j].CreatedAt)
	})

	if query.Since.Valid {
		var start int

		for start = 0; start < len(tmp); start++ {
			if query.Inclusive {
				if !tmp[start].CreatedAt.Before(query.Since.Time) {
					break
				}
			} else {
				if tmp[start].CreatedAt.After(query.Since.Time) {
					break
				}
			}
		}

		if start == len(tmp) {
			tmp = make([]*model.Message, 0)
		} else {
			tmp = tmp[start:]
		}
	}
	if query.Until.Valid {
		var end int

		for end = len(tmp) - 1; end >= 0; end-- {
			if query.Inclusive {
				if !tmp[end].CreatedAt.After(query.Until.Time) {
					break
				}
			} else {
				if tmp[end].CreatedAt.Before(query.Until.Time) {
					break
				}
			}
		}

		if end < 0 {
			tmp = make([]*model.Message, 0)
		} else {
			tmp = tmp[:end+1]
		}
	}

	if query.Offset < 0 {
		query.Offset = 0
	}

	if query.Limit <= 0 {
		query.Limit = math.MaxInt32
	}

	more = len(tmp) > query.Offset+query.Limit
	messages = make([]*model.Message, 0)
	for i := query.Offset; i < len(tmp) && i < query.Offset+query.Limit; i++ {
		messages = append(messages, tmp[i])
	}
	return
}

func (repo *TestRepository) GetArchivedMessagesByID(messageID uuid.UUID) ([]*model.ArchivedMessage, error) {
	panic("implement me")
}

func (repo *TestRepository) SetMessageUnread(userID, messageID uuid.UUID, noticeable bool) error {
	if userID == uuid.Nil || messageID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.MessageUnreadsLock.Lock()
	mMap, ok := repo.MessageUnreads[userID]
	if !ok {
		mMap = make(map[uuid.UUID]bool)
	}
	mMap[messageID] = true
	repo.MessageUnreads[userID] = mMap
	repo.MessageUnreadsLock.Unlock()
	return nil
}

func (repo *TestRepository) GetUnreadMessagesByUserID(userID uuid.UUID) ([]*model.Message, error) {
	result := make([]*model.Message, 0)
	repo.MessageUnreadsLock.RLock()
	repo.MessagesLock.RLock()
	for uid, mMap := range repo.MessageUnreads {
		if uid != userID {
			continue
		}
		for mid := range mMap {
			m, ok := repo.Messages[mid]
			if ok {
				result = append(result, &m)
			}
		}
	}
	repo.MessagesLock.RUnlock()
	repo.MessageUnreadsLock.RUnlock()
	sort.Slice(result, func(i, j int) bool {
		return result[j].CreatedAt.After(result[i].CreatedAt)
	})
	return result, nil
}

func (repo *TestRepository) DeleteUnreadsByChannelID(channelID, userID uuid.UUID) error {
	if channelID == uuid.Nil || userID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.MessageUnreadsLock.Lock()
	repo.MessagesLock.RLock()
	for uid, mMap := range repo.MessageUnreads {
		if uid != userID {
			continue
		}
		var deleted []uuid.UUID
		for mid := range mMap {
			m, ok := repo.Messages[mid]
			if ok {
				if m.ChannelID == channelID {
					deleted = append(deleted, mid)
				}
			}
		}
		for _, v := range deleted {
			delete(mMap, v)
		}
	}
	repo.MessagesLock.RUnlock()
	repo.MessageUnreadsLock.Unlock()
	return nil
}

func (repo *TestRepository) GetChannelLatestMessagesByUserID(userID uuid.UUID, limit int, subscribeOnly bool) ([]*model.Message, error) {
	panic("implement me")
}

func (repo *TestRepository) CreateMessageReport(messageID, reporterID uuid.UUID, reason string) error {
	if messageID == uuid.Nil || reporterID == uuid.Nil {
		return repository.ErrNilID
	}

	// make report
	report := model.MessageReport{
		ID:        uuid.Must(uuid.NewV4()),
		MessageID: messageID,
		Reporter:  reporterID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}
	repo.MessageReportsLock.Lock()
	defer repo.MessageReportsLock.Unlock()
	for _, v := range repo.MessageReports {
		if v.MessageID == messageID && v.Reporter == reporterID {
			return repository.ErrAlreadyExists
		}
	}
	repo.MessageReports = append(repo.MessageReports, report)
	return nil
}

func (repo *TestRepository) GetMessageReports(offset, limit int) ([]*model.MessageReport, error) {
	repo.MessageReportsLock.RLock()
	l := len(repo.MessageReports)
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = math.MaxInt32
	}
	result := make([]*model.MessageReport, 0)
	for i := offset; i < l && i < offset+limit; i++ {
		re := repo.MessageReports[i]
		result = append(result, &re)
	}
	repo.MessageReportsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) GetMessageReportsByMessageID(messageID uuid.UUID) ([]*model.MessageReport, error) {
	repo.MessageReportsLock.RLock()
	result := make([]*model.MessageReport, 0)
	for _, v := range repo.MessageReports {
		if v.MessageID == messageID {
			v := v
			result = append(result, &v)
		}
	}
	repo.MessageReportsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) GetMessageReportsByReporterID(reporterID uuid.UUID) ([]*model.MessageReport, error) {
	repo.MessageReportsLock.RLock()
	result := make([]*model.MessageReport, 0)
	for _, v := range repo.MessageReports {
		if v.Reporter == reporterID {
			v := v
			result = append(result, &v)
		}
	}
	repo.MessageReportsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) AddStampToMessage(messageID, stampID, userID uuid.UUID, count int) (ms *model.MessageStamp, err error) {
	panic("implement me")
}

func (repo *TestRepository) RemoveStampFromMessage(messageID, stampID, userID uuid.UUID) (err error) {
	panic("implement me")
}

func (repo *TestRepository) GetUserStampHistory(userID uuid.UUID, limit int) (h []*repository.UserStampHistory, err error) {
	panic("implement me")
}

func (repo *TestRepository) CreateStamp(name string, fileID, userID uuid.UUID) (s *model.Stamp, err error) {
	stamp := &model.Stamp{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      name,
		FileID:    fileID,
		CreatorID: userID, // uuid.Nilを許容する
	}

	repo.StampsLock.Lock()
	repo.FilesLock.RLock()
	defer repo.StampsLock.Unlock()
	defer repo.FilesLock.RUnlock()

	// 名前チェック
	if err := vd.Validate(name, validator.StampNameRuleRequired...); err != nil {
		return nil, repository.ArgError("name", "Name must be 1-32 characters of a-zA-Z0-9_-")
	}
	// 名前重複チェック
	for _, v := range repo.Stamps {
		if v.Name == name {
			return nil, repository.ErrAlreadyExists
		}
	}
	// ファイル存在チェック
	if fileID == uuid.Nil {
		return nil, repository.ArgError("fileID", "FileID's file is not found")
	}
	if _, ok := repo.Files[fileID]; !ok {
		return nil, repository.ArgError("fileID", "fileID's file is not found")
	}

	repo.Stamps[stamp.ID] = *stamp
	return stamp, nil
}

func (repo *TestRepository) UpdateStamp(id uuid.UUID, args repository.UpdateStampArgs) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}

	repo.StampsLock.Lock()
	repo.FilesLock.RLock()
	defer repo.StampsLock.Unlock()
	defer repo.FilesLock.RUnlock()

	s, ok := repo.Stamps[id]
	if !ok {
		return repository.ErrNotFound
	}

	if args.Name.Valid {
		if err := vd.Validate(args.Name.String, validator.StampNameRuleRequired...); err != nil {
			return repository.ArgError("args.Name", "Name must be 1-32 characters of a-zA-Z0-9_-")
		}

		// 重複チェック
		for _, v := range repo.Stamps {
			if v.Name == args.Name.String {
				return repository.ErrAlreadyExists
			}
		}
		s.Name = args.Name.String
	}
	if args.FileID.Valid {
		// 存在チェック
		if args.FileID.UUID == uuid.Nil {
			return repository.ArgError("args.FileID", "FileID's file is not found")
		}
		if _, ok := repo.Files[args.FileID.UUID]; !ok {
			return repository.ArgError("fileID", "fileID's file is not found")
		}

		s.FileID = args.FileID.UUID
	}
	if args.CreatorID.Valid {
		// uuid.Nilを許容する
		s.CreatorID = args.CreatorID.UUID
	}

	s.UpdatedAt = time.Now()
	repo.Stamps[id] = s
	return nil
}

func (repo *TestRepository) GetStamp(id uuid.UUID) (*model.Stamp, error) {
	if id == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	repo.StampsLock.RLock()
	s, ok := repo.Stamps[id]
	repo.StampsLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &s, nil
}

func (repo *TestRepository) DeleteStamp(id uuid.UUID) (err error) {
	if id == uuid.Nil {
		return repository.ErrNilID
	}
	repo.StampsLock.Lock()
	defer repo.StampsLock.Unlock()
	if _, ok := repo.Stamps[id]; !ok {
		return repository.ErrNotFound
	}
	delete(repo.Stamps, id)
	return nil
}

func (repo *TestRepository) GetAllStamps(excludeUnicode bool) (stamps []*model.Stamp, err error) {
	repo.StampsLock.RLock()
	for _, v := range repo.Stamps {
		v := v
		if excludeUnicode && v.IsUnicode {
			continue
		}
		stamps = append(stamps, &v)
	}
	repo.StampsLock.RUnlock()
	return
}

func (repo *TestRepository) StampExists(id uuid.UUID) (bool, error) {
	if id == uuid.Nil {
		return false, nil
	}
	repo.StampsLock.RLock()
	_, ok := repo.Stamps[id]
	repo.StampsLock.RUnlock()
	return ok, nil
}

func (repo *TestRepository) StampNameExists(name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	repo.StampsLock.RUnlock()
	defer repo.StampsLock.RUnlock()
	for _, v := range repo.Stamps {
		if v.Name == name {
			return true, nil
		}
	}
	return false, nil
}

func (repo *TestRepository) AddStar(userID, channelID uuid.UUID) error {
	if userID == uuid.Nil || channelID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.StarsLock.Lock()
	chMap, ok := repo.Stars[userID]
	if !ok {
		chMap = make(map[uuid.UUID]bool)
	}
	chMap[channelID] = true
	repo.Stars[userID] = chMap
	repo.StarsLock.Unlock()
	return nil
}

func (repo *TestRepository) RemoveStar(userID, channelID uuid.UUID) error {
	if userID == uuid.Nil || channelID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.StarsLock.Lock()
	chMap, ok := repo.Stars[userID]
	if ok {
		delete(chMap, channelID)
		repo.Stars[userID] = chMap
	}
	repo.StarsLock.Unlock()
	return nil
}

func (repo *TestRepository) GetStaredChannels(userID uuid.UUID) ([]uuid.UUID, error) {
	repo.StarsLock.RLock()
	result := make([]uuid.UUID, 0)
	chMap, ok := repo.Stars[userID]
	if ok {
		for id := range chMap {
			result = append(result, id)
		}
	}
	repo.StarsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) CreatePin(messageID, userID uuid.UUID) (*model.Pin, error) {
	if messageID == uuid.Nil || userID == uuid.Nil {
		return nil, repository.ErrNilID
	}
	repo.PinsLock.Lock()
	defer repo.PinsLock.Unlock()
	for _, pin := range repo.Pins {
		pin := pin
		if pin.MessageID == messageID {
			return &pin, nil
		}
	}
	p := model.Pin{
		ID:        uuid.Must(uuid.NewV4()),
		MessageID: messageID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	repo.Pins[p.ID] = p
	return &p, nil
}

func (repo *TestRepository) GetPin(id uuid.UUID) (*model.Pin, error) {
	repo.PinsLock.RLock()
	pin, ok := repo.Pins[id]
	repo.PinsLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	repo.MessagesLock.RLock()
	pin.Message = repo.Messages[pin.MessageID]
	repo.MessagesLock.RUnlock()
	return &pin, nil
}

func (repo *TestRepository) DeletePin(id uuid.UUID, userID uuid.UUID) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}
	repo.PinsLock.Lock()
	delete(repo.Pins, id)
	repo.PinsLock.Unlock()
	return nil
}

func (repo *TestRepository) GetPinsByChannelID(channelID uuid.UUID) ([]*model.Pin, error) {
	result := make([]*model.Pin, 0)
	repo.PinsLock.RLock()
	repo.MessagesLock.RLock()
	for _, p := range repo.Pins {
		m, ok := repo.Messages[p.MessageID]
		if ok && m.ChannelID == channelID {
			p := p
			p.Message = m
			result = append(result, &p)
		}
	}
	repo.MessagesLock.RUnlock()
	repo.PinsLock.RUnlock()
	return result, nil
}

func (repo *TestRepository) RegisterDevice(userID uuid.UUID, token string) (*model.Device, error) {
	panic("implement me")
}

func (repo *TestRepository) DeleteDeviceTokens(tokens []string) error {
	panic("implement me")
}

func (repo *TestRepository) GetDevicesByUserID(user uuid.UUID) (result []*model.Device, err error) {
	panic("implement me")
}

func (repo *TestRepository) GetDeviceTokens(users set.UUIDSet) (tokens []string, err error) {
	panic("implement me")
}

func (repo *TestRepository) GetAllDevices() (result []*model.Device, err error) {
	panic("implement me")
}

func (repo *TestRepository) GetAllDeviceTokens() (result []string, err error) {
	panic("implement me")
}

func (repo *TestRepository) GetFileMeta(fileID uuid.UUID) (model.FileMeta, error) {
	if fileID == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	repo.FilesLock.RLock()
	meta, ok := repo.Files[fileID]
	repo.FilesLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &fileMetaImpl{meta: &meta, fs: repo.FS}, nil
}

func (repo *TestRepository) DeleteFile(fileID uuid.UUID) error {
	if fileID == uuid.Nil {
		return repository.ErrNilID
	}
	repo.FilesLock.Lock()
	defer repo.FilesLock.Unlock()
	meta, ok := repo.Files[fileID]
	if !ok {
		return repository.ErrNotFound
	}
	delete(repo.Files, fileID)
	return repo.FS.DeleteByKey(meta.ID.String(), meta.Type)
}

func (repo *TestRepository) SaveFile(args repository.SaveFileArgs) (model.FileMeta, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	f := &model.File{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      args.FileName,
		Size:      args.FileSize,
		Mime:      args.MimeType,
		Type:      args.FileType,
		CreatorID: args.CreatorID,
		ChannelID: args.ChannelID,
		CreatedAt: time.Now(),
	}

	eg, ctx := errgroup.WithContext(context.Background())

	fileSrc, fileWriter := io.Pipe()
	thumbSrc, thumbWriter := io.Pipe()
	hash := md5.New()

	go func() {
		defer fileWriter.Close()
		defer thumbWriter.Close()
		_, _ = io.Copy(ioext.MultiWriter(fileWriter, hash, thumbWriter), args.Src) // 並列化してるけど、pipeじゃなくてbuffer使わないとpipeがブロックしてて意味無い疑惑
	}()

	// fileの保存
	eg.Go(func() error {
		defer fileSrc.Close()
		if err := repo.FS.SaveByKey(fileSrc, f.ID.String(), f.Name, f.Mime, f.Type); err != nil {
			return err
		}
		return nil
	})

	// サムネイルの生成
	eg.Go(func() error {
		// アップロードされたファイルの拡張子が間違えてたり、変なの送ってきた場合
		// サムネイルを生成しないだけで全体のエラーにはしない
		defer thumbSrc.Close()
		size, _ := repo.generateThumbnail(ctx, f, thumbSrc)
		if !size.Empty() {
			f.HasThumbnail = true
			f.ThumbnailMime = null.StringFrom("image/png")
			f.ThumbnailWidth = size.Size().X
			f.ThumbnailHeight = size.Size().Y
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	f.Hash = hex.EncodeToString(hash.Sum(nil))
	repo.FilesLock.Lock()
	repo.FilesACLLock.Lock()
	repo.Files[f.ID] = *f
	repo.FilesACL[f.ID] = args.ACL
	repo.FilesACLLock.Unlock()
	repo.FilesLock.Unlock()
	return &fileMetaImpl{meta: f, fs: repo.FS}, nil
}

func (repo *TestRepository) IsFileAccessible(fileID, userID uuid.UUID) (bool, error) {
	if fileID == uuid.Nil {
		return false, repository.ErrNilID
	}
	repo.FilesLock.RLock()
	_, ok := repo.Files[fileID]
	repo.FilesLock.RUnlock()
	if !ok {
		return false, repository.ErrNotFound
	}

	var allow bool
	repo.FilesACLLock.RLock()
	defer repo.FilesACLLock.RUnlock()
	for uid, a := range repo.FilesACL[fileID] {
		if uid == uuid.Nil || uid == userID {
			if a {
				allow = true
			} else {
				return false, nil
			}
		}
	}
	return allow, nil
}

func (repo *TestRepository) generateThumbnail(ctx context.Context, f *model.File, src io.Reader) (image.Rectangle, error) {
	orig, err := imaging.Decode(src, imaging.AutoOrientation(true))
	if err != nil {
		return image.Rectangle{}, err
	}

	img := imaging.Fit(orig, 360, 480, imaging.Linear)

	r, w := io.Pipe()
	go func() {
		_ = imaging.Encode(w, img, imaging.PNG)
		_ = w.Close()
	}()

	key := f.ID.String() + "-thumb"
	if err := repo.FS.SaveByKey(r, key, key+".png", "image/png", model.FileTypeThumbnail); err != nil {
		return image.Rectangle{}, err
	}
	return img.Bounds(), nil
}

func (repo *TestRepository) CreateWebhook(name, description string, channelID, creatorID uuid.UUID, secret string) (model.Webhook, error) {
	if len(name) == 0 || utf8.RuneCountInString(name) > 32 {
		return nil, repository.ArgError("name", "Name must be non-empty and shorter than 33 characters")
	}
	uid := uuid.Must(uuid.NewV4())
	bid := uuid.Must(uuid.NewV4())
	iconID, err := repository.GenerateIconFile(repo, name)
	if err != nil {
		return nil, err
	}

	u := model.User{
		ID:          uid,
		Name:        "Webhook#" + base64.RawStdEncoding.EncodeToString(uid.Bytes()),
		DisplayName: name,
		Icon:        iconID,
		Bot:         true,
		Status:      model.UserAccountStatusActive,
		Role:        role.Bot,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Profile: &model.UserProfile{
			UserID:    uid,
			UpdatedAt: time.Now(),
		},
	}
	wb := model.WebhookBot{
		ID:          bid,
		BotUserID:   uid,
		Description: description,
		Secret:      secret,
		ChannelID:   channelID,
		CreatorID:   creatorID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repo.WebhooksLock.Lock()
	repo.UsersLock.Lock()
	repo.ChannelsLock.RLock()
	defer repo.UsersLock.Unlock()
	defer repo.WebhooksLock.Unlock()
	defer repo.ChannelsLock.RUnlock()

	ch, ok := repo.Channels[channelID]
	if !ok {
		return nil, repository.ArgError("channelID", "the Channel is not found")
	}
	if !ch.IsPublic {
		return nil, repository.ArgError("channelID", "private channels are not allowed")
	}

	repo.Users[uid] = u
	repo.Webhooks[bid] = wb

	wb.BotUser = u
	return &wb, nil
}

func (repo *TestRepository) UpdateWebhook(id uuid.UUID, args repository.UpdateWebhookArgs) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}

	repo.WebhooksLock.Lock()
	repo.UsersLock.Lock()
	repo.ChannelsLock.RLock()
	defer repo.WebhooksLock.Unlock()
	defer repo.UsersLock.Unlock()
	defer repo.ChannelsLock.RUnlock()

	wb, ok := repo.Webhooks[id]
	if !ok {
		return repository.ErrNotFound
	}
	u := repo.Users[wb.GetBotUserID()]

	if args.Description.Valid {
		wb.Description = args.Description.String
		wb.UpdatedAt = time.Now()
	}
	if args.ChannelID.Valid {
		ch, ok := repo.Channels[args.ChannelID.UUID]
		if !ok {
			return repository.ArgError("args.ChannelID", "the Channel is not found")
		}
		if !ch.IsPublic {
			return repository.ArgError("args.ChannelID", "private channels are not allowed")
		}
		wb.ChannelID = args.ChannelID.UUID
		wb.UpdatedAt = time.Now()
	}
	if args.Secret.Valid {
		wb.Secret = args.Secret.String
		wb.UpdatedAt = time.Now()
	}
	if args.Name.Valid {
		if len(args.Name.String) == 0 || utf8.RuneCountInString(args.Name.String) > 32 {
			return repository.ArgError("args.Name", "Name must be non-empty and shorter than 33 characters")
		}
		u.DisplayName = args.Name.String
		u.UpdatedAt = time.Now()
	}

	repo.Webhooks[id] = wb
	repo.Users[u.ID] = u
	return nil
}

func (repo *TestRepository) DeleteWebhook(id uuid.UUID) error {
	if id == uuid.Nil {
		return repository.ErrNilID
	}
	repo.WebhooksLock.Lock()
	repo.UsersLock.Lock()
	defer repo.WebhooksLock.Unlock()
	defer repo.UsersLock.Unlock()
	wb, ok := repo.Webhooks[id]
	if !ok {
		return repository.ErrNotFound
	}
	delete(repo.Webhooks, id)
	u := repo.Users[wb.BotUserID]
	u.Status = model.UserAccountStatusDeactivated
	u.UpdatedAt = time.Now()
	repo.Users[wb.BotUserID] = u
	return nil
}

func (repo *TestRepository) GetWebhook(id uuid.UUID) (model.Webhook, error) {
	if id == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	repo.WebhooksLock.RLock()
	repo.UsersLock.RLock()
	defer repo.WebhooksLock.RUnlock()
	defer repo.UsersLock.RUnlock()
	w, ok := repo.Webhooks[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	w.BotUser = repo.Users[w.BotUserID]
	return &w, nil
}

func (repo *TestRepository) GetWebhookByBotUserID(id uuid.UUID) (model.Webhook, error) {
	if id == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	repo.WebhooksLock.RLock()
	repo.UsersLock.RLock()
	defer repo.WebhooksLock.RUnlock()
	defer repo.UsersLock.RUnlock()
	var (
		w  model.WebhookBot
		ok bool
	)
	for _, v := range repo.Webhooks {
		if v.BotUserID == id {
			w = v
			ok = true
			break
		}
	}
	if !ok {
		return nil, repository.ErrNotFound
	}
	w.BotUser = repo.Users[w.BotUserID]
	return &w, nil
}

func (repo *TestRepository) GetAllWebhooks() ([]model.Webhook, error) {
	arr := make([]model.Webhook, 0)
	repo.WebhooksLock.RLock()
	repo.UsersLock.RLock()
	for _, v := range repo.Webhooks {
		v := v
		v.BotUser = repo.Users[v.BotUserID]
		arr = append(arr, &v)
	}
	repo.UsersLock.RUnlock()
	repo.WebhooksLock.RUnlock()
	return arr, nil
}

func (repo *TestRepository) GetWebhooksByCreator(creatorID uuid.UUID) ([]model.Webhook, error) {
	arr := make([]model.Webhook, 0)
	if creatorID == uuid.Nil {
		return arr, nil
	}
	repo.WebhooksLock.RLock()
	repo.UsersLock.RLock()
	for _, v := range repo.Webhooks {
		if v.CreatorID == creatorID {
			v := v
			v.BotUser = repo.Users[v.BotUserID]
			arr = append(arr, &v)
		}
	}
	repo.UsersLock.RUnlock()
	repo.WebhooksLock.RUnlock()
	return arr, nil
}

func (repo *TestRepository) GetClient(id string) (*model.OAuth2Client, error) {
	if len(id) == 0 {
		return nil, repository.ErrNotFound
	}
	repo.OAuth2ClientsLock.RLock()
	c, ok := repo.OAuth2Clients[id]
	repo.OAuth2ClientsLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &c, nil
}

func (repo *TestRepository) GetClients(query repository.GetClientsQuery) ([]*model.OAuth2Client, error) {
	panic("implement me")
}

func (repo *TestRepository) SaveClient(client *model.OAuth2Client) error {
	repo.OAuth2ClientsLock.Lock()
	client.CreatedAt = time.Now()
	client.UpdatedAt = time.Now()
	repo.OAuth2Clients[client.ID] = *client
	repo.OAuth2ClientsLock.Unlock()
	return nil
}

func (repo *TestRepository) UpdateClient(clientID string, args repository.UpdateClientArgs) error {
	panic("implement me")
}

func (repo *TestRepository) DeleteClient(id string) error {
	if len(id) == 0 {
		return nil
	}
	repo.OAuth2ClientsLock.Lock()
	repo.OAuth2AuthorizesLock.Lock()
	repo.OAuth2TokensLock.Lock()
	targetT := make([]uuid.UUID, 0)
	for k, v := range repo.OAuth2Tokens {
		if v.ClientID == id {
			targetT = append(targetT, k)
		}
	}
	for _, v := range targetT {
		delete(repo.OAuth2Tokens, v)
	}
	targetA := make([]string, 0)
	for k, v := range repo.OAuth2Authorizes {
		if v.ClientID == id {
			targetA = append(targetA, k)
		}
	}
	for _, v := range targetA {
		delete(repo.OAuth2Authorizes, v)
	}
	delete(repo.OAuth2Clients, id)
	repo.OAuth2TokensLock.Unlock()
	repo.OAuth2AuthorizesLock.Unlock()
	repo.OAuth2ClientsLock.Unlock()
	return nil
}

func (repo *TestRepository) SaveAuthorize(data *model.OAuth2Authorize) error {
	repo.OAuth2AuthorizesLock.Lock()
	data.CreatedAt = time.Now()
	repo.OAuth2Authorizes[data.Code] = *data
	repo.OAuth2AuthorizesLock.Unlock()
	return nil
}

func (repo *TestRepository) GetAuthorize(code string) (*model.OAuth2Authorize, error) {
	if len(code) == 0 {
		return nil, repository.ErrNotFound
	}
	repo.OAuth2AuthorizesLock.RLock()
	a, ok := repo.OAuth2Authorizes[code]
	repo.OAuth2AuthorizesLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &a, nil
}

func (repo *TestRepository) DeleteAuthorize(code string) error {
	if len(code) == 0 {
		return nil
	}
	repo.OAuth2AuthorizesLock.Lock()
	delete(repo.OAuth2Authorizes, code)
	repo.OAuth2AuthorizesLock.Unlock()
	return nil
}

func (repo *TestRepository) IssueToken(client *model.OAuth2Client, userID uuid.UUID, redirectURI string, scope model.AccessScopes, expire int, refresh bool) (*model.OAuth2Token, error) {
	newToken := &model.OAuth2Token{
		ID:             uuid.Must(uuid.NewV4()),
		UserID:         userID,
		RedirectURI:    redirectURI,
		AccessToken:    utils.RandAlphabetAndNumberString(36),
		RefreshToken:   utils.RandAlphabetAndNumberString(36),
		RefreshEnabled: refresh,
		CreatedAt:      time.Now(),
		ExpiresIn:      expire,
		Scopes:         scope,
	}

	if client != nil {
		newToken.ClientID = client.ID
	}

	repo.OAuth2TokensLock.Lock()
	repo.OAuth2Tokens[newToken.ID] = *newToken
	repo.OAuth2TokensLock.Unlock()
	return newToken, nil
}

func (repo *TestRepository) GetTokenByID(id uuid.UUID) (*model.OAuth2Token, error) {
	if id == uuid.Nil {
		return nil, repository.ErrNotFound
	}
	repo.OAuth2TokensLock.RLock()
	t, ok := repo.OAuth2Tokens[id]
	repo.OAuth2TokensLock.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return &t, nil
}

func (repo *TestRepository) DeleteTokenByID(id uuid.UUID) error {
	if id == uuid.Nil {
		return nil
	}
	repo.OAuth2TokensLock.Lock()
	delete(repo.OAuth2Tokens, id)
	repo.OAuth2TokensLock.Unlock()
	return nil
}

func (repo *TestRepository) GetTokenByAccess(access string) (*model.OAuth2Token, error) {
	if len(access) == 0 {
		return nil, repository.ErrNotFound
	}
	repo.OAuth2TokensLock.RLock()
	defer repo.OAuth2TokensLock.RUnlock()
	for _, v := range repo.OAuth2Tokens {
		if v.AccessToken == access {
			return &v, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (repo *TestRepository) DeleteTokenByAccess(access string) error {
	if len(access) == 0 {
		return nil
	}
	repo.OAuth2TokensLock.Lock()
	defer repo.OAuth2TokensLock.Unlock()
	for k, v := range repo.OAuth2Tokens {
		if v.AccessToken == access {
			delete(repo.OAuth2Tokens, k)
			return nil
		}
	}
	return nil
}

func (repo *TestRepository) GetTokenByRefresh(refresh string) (*model.OAuth2Token, error) {
	if len(refresh) == 0 {
		return nil, repository.ErrNotFound
	}
	repo.OAuth2TokensLock.RLock()
	defer repo.OAuth2TokensLock.RUnlock()
	for _, v := range repo.OAuth2Tokens {
		if v.RefreshEnabled && v.RefreshToken == refresh {
			return &v, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (repo *TestRepository) DeleteTokenByRefresh(refresh string) error {
	if len(refresh) == 0 {
		return nil
	}
	repo.OAuth2TokensLock.Lock()
	defer repo.OAuth2TokensLock.Unlock()
	for k, v := range repo.OAuth2Tokens {
		if v.RefreshEnabled && v.RefreshToken == refresh {
			delete(repo.OAuth2Tokens, k)
			return nil
		}
	}
	return nil
}

func (repo *TestRepository) GetTokensByUser(userID uuid.UUID) ([]*model.OAuth2Token, error) {
	ts := make([]*model.OAuth2Token, 0)
	if userID == uuid.Nil {
		return ts, nil
	}
	repo.OAuth2TokensLock.RLock()
	for _, v := range repo.OAuth2Tokens {
		v := v
		if v.UserID == userID {
			ts = append(ts, &v)
		}
	}
	repo.OAuth2TokensLock.RUnlock()
	return ts, nil
}

func (repo *TestRepository) DeleteTokenByUser(userID uuid.UUID) error {
	if userID == uuid.Nil {
		return nil
	}
	repo.OAuth2TokensLock.Lock()
	target := make([]uuid.UUID, 0)
	for k, v := range repo.OAuth2Tokens {
		if v.UserID == userID {
			target = append(target, k)
		}
	}
	for _, v := range target {
		delete(repo.OAuth2Tokens, v)
	}
	repo.OAuth2TokensLock.Unlock()
	return nil
}

func (repo *TestRepository) DeleteTokenByClient(clientID string) error {
	if len(clientID) == 0 {
		return nil
	}
	repo.OAuth2TokensLock.Lock()
	target := make([]uuid.UUID, 0)
	for k, v := range repo.OAuth2Tokens {
		if v.ClientID == clientID {
			target = append(target, k)
		}
	}
	for _, v := range target {
		delete(repo.OAuth2Tokens, v)
	}
	repo.OAuth2TokensLock.Unlock()
	return nil
}

func (repo *TestRepository) CreateBot(name, displayName, description string, creatorID uuid.UUID, webhookURL string) (*model.Bot, error) {
	panic("implement me")
}

func (repo *TestRepository) GetBotByID(id uuid.UUID) (*model.Bot, error) {
	panic("implement me")
}

func (repo *TestRepository) GetBotByCode(code string) (*model.Bot, error) {
	panic("implement me")
}

func (repo *TestRepository) GetBots(query repository.BotsQuery) ([]*model.Bot, error) {
	panic("implement me")
}

func (repo *TestRepository) ChangeBotState(id uuid.UUID, state model.BotState) error {
	panic("implement me")
}

func (repo *TestRepository) DeleteBot(id uuid.UUID) error {
	panic("implement me")
}

func (repo *TestRepository) AddBotToChannel(botID, channelID uuid.UUID) error {
	panic("implement me")
}

func (repo *TestRepository) RemoveBotFromChannel(botID, channelID uuid.UUID) error {
	panic("implement me")
}

func (repo *TestRepository) GetParticipatingChannelIDsByBot(botID uuid.UUID) ([]uuid.UUID, error) {
	panic("implement me")
}

func (repo *TestRepository) CreateClipFolder(userID uuid.UUID, name, description string) (*model.ClipFolder, error) {
	panic("implement me")
}

func (repo *TestRepository) UpdateClipFolder(folderID uuid.UUID, name, description null.String) error {
	panic("implement me")
}

func (repo *TestRepository) DeleteClipFolder(folderID uuid.UUID) error {
	panic("implement me")
}

func (repo *TestRepository) DeleteClipFolderMessage(folderID, messageID uuid.UUID) error {
	panic("implement me")
}

func (repo *TestRepository) AddClipFolderMessage(folderID, messageID uuid.UUID) (*model.ClipFolderMessage, error) {
	panic("implement me")
}

func (repo *TestRepository) GetClipFoldersByUserID(userID uuid.UUID) ([]*model.ClipFolder, error) {
	panic("implement me")
}

func (repo *TestRepository) GetClipFolder(folderID uuid.UUID) (*model.ClipFolder, error) {
	panic("implement me")
}

func (repo *TestRepository) GetClipFolderMessages(folderID uuid.UUID, query repository.ClipFolderMessageQuery) (messages []*model.ClipFolderMessage, more bool, err error) {
	panic("implement me")
}

type fileMetaImpl struct {
	meta *model.File
	fs   storage.FileStorage
}

func (f *fileMetaImpl) GetID() uuid.UUID {
	return f.meta.ID
}

func (f *fileMetaImpl) GetFileName() string {
	return f.meta.Name
}

func (f *fileMetaImpl) GetMIMEType() string {
	return f.meta.Mime
}

func (f *fileMetaImpl) GetFileSize() int64 {
	return f.meta.Size
}

func (f *fileMetaImpl) GetFileType() string {
	return f.meta.Type
}

func (f *fileMetaImpl) GetCreatorID() uuid.NullUUID {
	return f.meta.CreatorID
}

func (f *fileMetaImpl) GetMD5Hash() string {
	return f.meta.Hash
}

func (f *fileMetaImpl) HasThumbnail() bool {
	return f.meta.HasThumbnail
}

func (f *fileMetaImpl) GetThumbnailMIMEType() string {
	return f.meta.ThumbnailMime.String
}

func (f *fileMetaImpl) GetThumbnailWidth() int {
	return f.meta.ThumbnailWidth
}

func (f *fileMetaImpl) GetThumbnailHeight() int {
	return f.meta.ThumbnailHeight
}

func (f *fileMetaImpl) GetUploadChannelID() uuid.NullUUID {
	return f.meta.ChannelID
}

func (f *fileMetaImpl) GetCreatedAt() time.Time {
	return f.meta.CreatedAt
}

func (f *fileMetaImpl) Open() (ioext.ReadSeekCloser, error) {
	return f.fs.OpenFileByKey(f.GetID().String(), f.GetFileType())
}

func (f *fileMetaImpl) OpenThumbnail() (ioext.ReadSeekCloser, error) {
	if !f.HasThumbnail() {
		return nil, repository.ErrNotFound
	}
	return f.fs.OpenFileByKey(f.GetID().String()+"-thumb", model.FileTypeThumbnail)
}

func (f *fileMetaImpl) GetAlternativeURL() string {
	url, _ := f.fs.GenerateAccessURL(f.GetID().String(), f.GetFileType())
	return url
}

func (repo *TestRepository) GetFiles(q repository.FilesQuery) (result []model.FileMeta, more bool, err error) {
	panic("implement me")
}
