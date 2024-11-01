import {
  Checkbox,
  IconButton,
  ListItem,
  ListItemSecondaryAction,
  ListItemText
} from '@material-ui/core'
import DeleteIcon from '@material-ui/icons/Delete'
import EditIcon from '@material-ui/icons/Edit'
import React, { memo, useContext } from 'react'
import EditTodoForm from '../EditTodoForm'
import { DispatchContext } from '../context/todos.context'
import useToggleState from '../hooks/useToggleState'

function Todo({ id, task, completed }) {
  const [isEditing, toggle] = useToggleState(false)
  const dispatch = useContext(DispatchContext)

  return (
    <ListItem style={{ height: '64px' }}>
      {isEditing ? (
        <EditTodoForm task={task} id={id} toggleEditing={toggle} />
      ) : (
        <>
          <Checkbox
            tabIndex={-1}
            checked={completed}
            onClick={() => dispatch({ type: 'TOGGLE', id: id })}
          />

          <ListItemText style={{ textDecoration: completed ? 'line-through' : 'none' }}>
            {task}
          </ListItemText>

          <ListItemSecondaryAction>
            <IconButton aria-label='Delete' onClick={() => dispatch({ type: 'REMOVE', id: id })}>
              <DeleteIcon />
            </IconButton>
            <IconButton aria-label='Edit' onClick={toggle}>
              <EditIcon />
            </IconButton>
          </ListItemSecondaryAction>
        </>
      )}
    </ListItem>
  )
}

export default memo(Todo)
