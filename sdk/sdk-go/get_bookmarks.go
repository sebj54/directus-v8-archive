/* 
 * directus.io
 *
 * API for directus.io
 *
 * OpenAPI spec version: 1.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package directussdk

type GetBookmarks struct {

	Meta *GetBookmarksMeta `json:"meta,omitempty"`

	Data []GetBookmarksData `json:"data,omitempty"`
}