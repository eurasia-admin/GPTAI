// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: chat_prompt.sql

package sqlc_queries

import (
	"context"
)

const createChatPrompt = `-- name: CreateChatPrompt :one
INSERT INTO chat_prompt (uuid, chat_session_uuid, role, content, token_count, user_id, created_by, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
`

type CreateChatPromptParams struct {
	Uuid            string
	ChatSessionUuid string
	Role            string
	Content         string
	TokenCount      int32
	UserID          int32
	CreatedBy       int32
	UpdatedBy       int32
}

func (q *Queries) CreateChatPrompt(ctx context.Context, arg CreateChatPromptParams) (ChatPrompt, error) {
	row := q.db.QueryRowContext(ctx, createChatPrompt,
		arg.Uuid,
		arg.ChatSessionUuid,
		arg.Role,
		arg.Content,
		arg.TokenCount,
		arg.UserID,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i ChatPrompt
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.ChatSessionUuid,
		&i.Role,
		&i.Content,
		&i.Score,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.IsDeleted,
		&i.TokenCount,
	)
	return i, err
}

const deleteChatPrompt = `-- name: DeleteChatPrompt :exec
UPDATE chat_prompt 
SET is_deleted = true, updated_at = now()
WHERE id = $1
`

func (q *Queries) DeleteChatPrompt(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteChatPrompt, id)
	return err
}

const deleteChatPromptByUUID = `-- name: DeleteChatPromptByUUID :exec
UPDATE chat_prompt
SET is_deleted = true, updated_at = now()
WHERE uuid = $1
`

func (q *Queries) DeleteChatPromptByUUID(ctx context.Context, uuid string) error {
	_, err := q.db.ExecContext(ctx, deleteChatPromptByUUID, uuid)
	return err
}

const getAllChatPrompts = `-- name: GetAllChatPrompts :many
SELECT id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count FROM chat_prompt 
WHERE is_deleted = false
ORDER BY id
`

func (q *Queries) GetAllChatPrompts(ctx context.Context) ([]ChatPrompt, error) {
	rows, err := q.db.QueryContext(ctx, getAllChatPrompts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatPrompt
	for rows.Next() {
		var i ChatPrompt
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.ChatSessionUuid,
			&i.Role,
			&i.Content,
			&i.Score,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.IsDeleted,
			&i.TokenCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChatPromptByID = `-- name: GetChatPromptByID :one
SELECT id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count FROM chat_prompt
WHERE is_deleted = false and  id = $1
`

func (q *Queries) GetChatPromptByID(ctx context.Context, id int32) (ChatPrompt, error) {
	row := q.db.QueryRowContext(ctx, getChatPromptByID, id)
	var i ChatPrompt
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.ChatSessionUuid,
		&i.Role,
		&i.Content,
		&i.Score,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.IsDeleted,
		&i.TokenCount,
	)
	return i, err
}

const getChatPromptsBySessionUUID = `-- name: GetChatPromptsBySessionUUID :many
SELECT id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
FROM chat_prompt 
WHERE chat_session_uuid = $1 and is_deleted = false
ORDER BY id
`

func (q *Queries) GetChatPromptsBySessionUUID(ctx context.Context, chatSessionUuid string) ([]ChatPrompt, error) {
	rows, err := q.db.QueryContext(ctx, getChatPromptsBySessionUUID, chatSessionUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatPrompt
	for rows.Next() {
		var i ChatPrompt
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.ChatSessionUuid,
			&i.Role,
			&i.Content,
			&i.Score,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.IsDeleted,
			&i.TokenCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChatPromptsByUserID = `-- name: GetChatPromptsByUserID :many
SELECT id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
FROM chat_prompt 
WHERE user_id = $1 and is_deleted = false
ORDER BY id
`

func (q *Queries) GetChatPromptsByUserID(ctx context.Context, userID int32) ([]ChatPrompt, error) {
	rows, err := q.db.QueryContext(ctx, getChatPromptsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatPrompt
	for rows.Next() {
		var i ChatPrompt
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.ChatSessionUuid,
			&i.Role,
			&i.Content,
			&i.Score,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.IsDeleted,
			&i.TokenCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChatPromptsBysession_uuid = `-- name: GetChatPromptsBysession_uuid :many
SELECT id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
FROM chat_prompt 
WHERE chat_session_uuid = $1 and is_deleted = false
ORDER BY id
`

func (q *Queries) GetChatPromptsBysession_uuid(ctx context.Context, chatSessionUuid string) ([]ChatPrompt, error) {
	rows, err := q.db.QueryContext(ctx, getChatPromptsBysession_uuid, chatSessionUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatPrompt
	for rows.Next() {
		var i ChatPrompt
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.ChatSessionUuid,
			&i.Role,
			&i.Content,
			&i.Score,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.IsDeleted,
			&i.TokenCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneChatPromptBySessionUUID = `-- name: GetOneChatPromptBySessionUUID :one
SELECT id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
FROM chat_prompt 
WHERE chat_session_uuid = $1 and is_deleted = false
ORDER BY id
LIMIT 1
`

func (q *Queries) GetOneChatPromptBySessionUUID(ctx context.Context, chatSessionUuid string) (ChatPrompt, error) {
	row := q.db.QueryRowContext(ctx, getOneChatPromptBySessionUUID, chatSessionUuid)
	var i ChatPrompt
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.ChatSessionUuid,
		&i.Role,
		&i.Content,
		&i.Score,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.IsDeleted,
		&i.TokenCount,
	)
	return i, err
}

const hasChatPromptPermission = `-- name: HasChatPromptPermission :one
SELECT COUNT(*) > 0 as has_permission
FROM chat_prompt cp
INNER JOIN auth_user au ON cp.user_id = au.id
WHERE cp.id = $1 AND (cp.user_id = $2 OR au.is_superuser) AND cp.is_deleted = false
`

type HasChatPromptPermissionParams struct {
	ID     int32
	UserID int32
}

func (q *Queries) HasChatPromptPermission(ctx context.Context, arg HasChatPromptPermissionParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, hasChatPromptPermission, arg.ID, arg.UserID)
	var has_permission bool
	err := row.Scan(&has_permission)
	return has_permission, err
}

const updateChatPrompt = `-- name: UpdateChatPrompt :one
UPDATE chat_prompt SET chat_session_uuid = $2, role = $3, content = $4, score = $5, user_id = $6, updated_at = now(), updated_by = $7
WHERE id = $1
RETURNING id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
`

type UpdateChatPromptParams struct {
	ID              int32
	ChatSessionUuid string
	Role            string
	Content         string
	Score           float64
	UserID          int32
	UpdatedBy       int32
}

func (q *Queries) UpdateChatPrompt(ctx context.Context, arg UpdateChatPromptParams) (ChatPrompt, error) {
	row := q.db.QueryRowContext(ctx, updateChatPrompt,
		arg.ID,
		arg.ChatSessionUuid,
		arg.Role,
		arg.Content,
		arg.Score,
		arg.UserID,
		arg.UpdatedBy,
	)
	var i ChatPrompt
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.ChatSessionUuid,
		&i.Role,
		&i.Content,
		&i.Score,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.IsDeleted,
		&i.TokenCount,
	)
	return i, err
}

const updateChatPromptByUUID = `-- name: UpdateChatPromptByUUID :one
UPDATE chat_prompt SET content = $2, token_count = $3, updated_at = now()
WHERE uuid = $1 and is_deleted = false
RETURNING id, uuid, chat_session_uuid, role, content, score, user_id, created_at, updated_at, created_by, updated_by, is_deleted, token_count
`

type UpdateChatPromptByUUIDParams struct {
	Uuid       string
	Content    string
	TokenCount int32
}

func (q *Queries) UpdateChatPromptByUUID(ctx context.Context, arg UpdateChatPromptByUUIDParams) (ChatPrompt, error) {
	row := q.db.QueryRowContext(ctx, updateChatPromptByUUID, arg.Uuid, arg.Content, arg.TokenCount)
	var i ChatPrompt
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.ChatSessionUuid,
		&i.Role,
		&i.Content,
		&i.Score,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.IsDeleted,
		&i.TokenCount,
	)
	return i, err
}
