package com.github.redxiiikk.codeReviewTool.plugins.jetbrains

import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.actionSystem.CommonDataKeys

class CreateCommentAction : AnAction() {
    override fun actionPerformed(action: AnActionEvent) {
        val editor = action.getData(CommonDataKeys.EDITOR)
        val selectModel = editor?.selectionModel

        if (selectModel != null) {
            if (selectModel.hasSelection()) {
                println("CreateCommentAction: single: ${selectModel.leadSelectionPosition?.line}")
            } else {
                print("CreateCommentAction: selected: ${selectModel.selectionStartPosition?.line} - ${selectModel.selectionEndPosition?.line}")
            }
        }

    }
}
