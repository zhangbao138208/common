// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.skig.tech/zero-core/common/ent/wingameslot"
)

// WinGameSlot is the model entity for the WinGameSlot schema.
type WinGameSlot struct {
	config `json:"-"`
	// ID of the ent.
	// ID(关联BrandGameId)
	ID string `json:"id,omitempty"`
	// 游戏ID(关联game_list)
	GameID int32 `json:"game_id,omitempty"`
	// 游戏大类类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能
	GameGroupID int8 `json:"game_group_id,omitempty"`
	// 游戏平台id
	PlatID int32 `json:"plat_id,omitempty"`
	// 游戏提供者
	Provider string `json:"provider,omitempty"`
	// 简体名称
	Name string `json:"name,omitempty"`
	// 游戏名字(中文)
	NameZh string `json:"name_zh,omitempty"`
	// 英文图片
	Img string `json:"img,omitempty"`
	// 新版游戏图片
	ImgNew string `json:"img_new,omitempty"`
	// 是否新游戏:1-是 0-否
	IsNew bool `json:"is_new,omitempty"`
	// 是否推荐主页 0否 1是
	IsCasino int8 `json:"is_casino,omitempty"`
	// 游戏类型ID(0r code)
	GameTypeID string `json:"game_type_id,omitempty"`
	// 游戏类型名称
	GameTypeName string `json:"game_type_name,omitempty"`
	// 收藏值
	FavoriteStar int32 `json:"favorite_star,omitempty"`
	// 热度
	HotStar int32 `json:"hot_star,omitempty"`
	// 排序
	Sort int32 `json:"sort,omitempty"`
	// 状态:1-启用 0-停用
	Status int8 `json:"status,omitempty"`
	// 设备:0-all 1-pc 2-h5
	Device int8 `json:"device,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int32 `json:"updated_at,omitempty"`
	// 最后更新人
	UpdatedUser string `json:"updated_user,omitempty"`
	// 维护信息
	Maintenance string `json:"maintenance,omitempty"`
	// 操作人姓名
	OperatorName string `json:"operator_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WinGameSlot) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case wingameslot.FieldIsNew:
			values[i] = new(sql.NullBool)
		case wingameslot.FieldGameID, wingameslot.FieldGameGroupID, wingameslot.FieldPlatID, wingameslot.FieldIsCasino, wingameslot.FieldFavoriteStar, wingameslot.FieldHotStar, wingameslot.FieldSort, wingameslot.FieldStatus, wingameslot.FieldDevice, wingameslot.FieldCreatedAt, wingameslot.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case wingameslot.FieldID, wingameslot.FieldProvider, wingameslot.FieldName, wingameslot.FieldNameZh, wingameslot.FieldImg, wingameslot.FieldImgNew, wingameslot.FieldGameTypeID, wingameslot.FieldGameTypeName, wingameslot.FieldUpdatedUser, wingameslot.FieldMaintenance, wingameslot.FieldOperatorName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WinGameSlot fields.
func (wgs *WinGameSlot) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wingameslot.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				wgs.ID = value.String
			}
		case wingameslot.FieldGameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_id", values[i])
			} else if value.Valid {
				wgs.GameID = int32(value.Int64)
			}
		case wingameslot.FieldGameGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field game_group_id", values[i])
			} else if value.Valid {
				wgs.GameGroupID = int8(value.Int64)
			}
		case wingameslot.FieldPlatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field plat_id", values[i])
			} else if value.Valid {
				wgs.PlatID = int32(value.Int64)
			}
		case wingameslot.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				wgs.Provider = value.String
			}
		case wingameslot.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wgs.Name = value.String
			}
		case wingameslot.FieldNameZh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_zh", values[i])
			} else if value.Valid {
				wgs.NameZh = value.String
			}
		case wingameslot.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				wgs.Img = value.String
			}
		case wingameslot.FieldImgNew:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img_new", values[i])
			} else if value.Valid {
				wgs.ImgNew = value.String
			}
		case wingameslot.FieldIsNew:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_new", values[i])
			} else if value.Valid {
				wgs.IsNew = value.Bool
			}
		case wingameslot.FieldIsCasino:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_casino", values[i])
			} else if value.Valid {
				wgs.IsCasino = int8(value.Int64)
			}
		case wingameslot.FieldGameTypeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_type_id", values[i])
			} else if value.Valid {
				wgs.GameTypeID = value.String
			}
		case wingameslot.FieldGameTypeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_type_name", values[i])
			} else if value.Valid {
				wgs.GameTypeName = value.String
			}
		case wingameslot.FieldFavoriteStar:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field favorite_star", values[i])
			} else if value.Valid {
				wgs.FavoriteStar = int32(value.Int64)
			}
		case wingameslot.FieldHotStar:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hot_star", values[i])
			} else if value.Valid {
				wgs.HotStar = int32(value.Int64)
			}
		case wingameslot.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				wgs.Sort = int32(value.Int64)
			}
		case wingameslot.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				wgs.Status = int8(value.Int64)
			}
		case wingameslot.FieldDevice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				wgs.Device = int8(value.Int64)
			}
		case wingameslot.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wgs.CreatedAt = int32(value.Int64)
			}
		case wingameslot.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wgs.UpdatedAt = int32(value.Int64)
			}
		case wingameslot.FieldUpdatedUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_user", values[i])
			} else if value.Valid {
				wgs.UpdatedUser = value.String
			}
		case wingameslot.FieldMaintenance:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field maintenance", values[i])
			} else if value.Valid {
				wgs.Maintenance = value.String
			}
		case wingameslot.FieldOperatorName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operator_name", values[i])
			} else if value.Valid {
				wgs.OperatorName = value.String
			}
		default:
			wgs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WinGameSlot.
// This includes values selected through modifiers, order, etc.
func (wgs *WinGameSlot) Value(name string) (ent.Value, error) {
	return wgs.selectValues.Get(name)
}

// Update returns a builder for updating this WinGameSlot.
// Note that you need to call WinGameSlot.Unwrap() before calling this method if this WinGameSlot
// was returned from a transaction, and the transaction was committed or rolled back.
func (wgs *WinGameSlot) Update() *WinGameSlotUpdateOne {
	return NewWinGameSlotClient(wgs.config).UpdateOne(wgs)
}

// Unwrap unwraps the WinGameSlot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wgs *WinGameSlot) Unwrap() *WinGameSlot {
	_tx, ok := wgs.config.driver.(*txDriver)
	if !ok {
		panic("ent: WinGameSlot is not a transactional entity")
	}
	wgs.config.driver = _tx.drv
	return wgs
}

// String implements the fmt.Stringer.
func (wgs *WinGameSlot) String() string {
	var builder strings.Builder
	builder.WriteString("WinGameSlot(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wgs.ID))
	builder.WriteString("game_id=")
	builder.WriteString(fmt.Sprintf("%v", wgs.GameID))
	builder.WriteString(", ")
	builder.WriteString("game_group_id=")
	builder.WriteString(fmt.Sprintf("%v", wgs.GameGroupID))
	builder.WriteString(", ")
	builder.WriteString("plat_id=")
	builder.WriteString(fmt.Sprintf("%v", wgs.PlatID))
	builder.WriteString(", ")
	builder.WriteString("provider=")
	builder.WriteString(wgs.Provider)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(wgs.Name)
	builder.WriteString(", ")
	builder.WriteString("name_zh=")
	builder.WriteString(wgs.NameZh)
	builder.WriteString(", ")
	builder.WriteString("img=")
	builder.WriteString(wgs.Img)
	builder.WriteString(", ")
	builder.WriteString("img_new=")
	builder.WriteString(wgs.ImgNew)
	builder.WriteString(", ")
	builder.WriteString("is_new=")
	builder.WriteString(fmt.Sprintf("%v", wgs.IsNew))
	builder.WriteString(", ")
	builder.WriteString("is_casino=")
	builder.WriteString(fmt.Sprintf("%v", wgs.IsCasino))
	builder.WriteString(", ")
	builder.WriteString("game_type_id=")
	builder.WriteString(wgs.GameTypeID)
	builder.WriteString(", ")
	builder.WriteString("game_type_name=")
	builder.WriteString(wgs.GameTypeName)
	builder.WriteString(", ")
	builder.WriteString("favorite_star=")
	builder.WriteString(fmt.Sprintf("%v", wgs.FavoriteStar))
	builder.WriteString(", ")
	builder.WriteString("hot_star=")
	builder.WriteString(fmt.Sprintf("%v", wgs.HotStar))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", wgs.Sort))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", wgs.Status))
	builder.WriteString(", ")
	builder.WriteString("device=")
	builder.WriteString(fmt.Sprintf("%v", wgs.Device))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", wgs.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", wgs.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_user=")
	builder.WriteString(wgs.UpdatedUser)
	builder.WriteString(", ")
	builder.WriteString("maintenance=")
	builder.WriteString(wgs.Maintenance)
	builder.WriteString(", ")
	builder.WriteString("operator_name=")
	builder.WriteString(wgs.OperatorName)
	builder.WriteByte(')')
	return builder.String()
}

// WinGameSlots is a parsable slice of WinGameSlot.
type WinGameSlots []*WinGameSlot
