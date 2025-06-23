package services

import "weeklytask-8/models"

func PaginateMenu(listMenu []models.ListMenu, page, itemsPerPage int) []models.ListMenu {
    start := page * itemsPerPage
    end := start + itemsPerPage
    if start > len(listMenu) {
        return []models.ListMenu{}
    }
    if end > len(listMenu) {
        end = len(listMenu)
    }
    return listMenu[start:end]
}