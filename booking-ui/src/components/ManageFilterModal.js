import React from 'react';
import { useForm, useFieldArray, Controller } from 'react-hook-form';
import Select from 'react-select';
import Modal from 'react-modal';

Modal.setAppElement('#root');

function ManageFilterModal({ isOpen, onRequestClose, fields = [], setFields, filterableFields }) {
  const { handleSubmit, control, register } = useForm({ defaultValues: { fields: fields } });
  const { fields: formFields, append, remove } = useFieldArray({
    control,
    name: 'fields',
  });

  const options = filterableFields.map((field) => ({ value: field, label: field }));

  const onSubmit = (data) => {
    setFields(data.fields);
    onRequestClose();
  };

  React.useEffect(() => {
    if (!isOpen) {
      handleSubmit(onSubmit)();
    }
  }, [isOpen]);

  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onRequestClose}
      contentLabel="Manage Filters"
      className="modal-content"
      overlayClassName="modal-overlay"
    >
      <form onSubmit={handleSubmit(onSubmit)}>
        {formFields.map((item, index) => (
          <div key={item.id} className="field-container">
            <button type="button" onClick={() => remove(index)} className="remove-field-btn">
              x
            </button>
            <div className="filter-field-container">
              <label className="dropdown-container">
                <Controller
                  name={`fields[${index}].field`}
                  control={control}
                  defaultValue={item.field}
                  render={({ field }) => (
                    <Select
                      {...field}
                      placeholder="field"
                      options={options}
                      className="filter-field-dropdown"
                      value={options.find((option) => option.value === field.value)}
                      onChange={(value) => field.onChange(value.value)}
                    />
                  )}
                />
              </label>
              <input {...register(`fields[${index}].value`)} placeholder="value" defaultValue={item.value} className="filter-field-input" />
            </div>
          </div>
        ))}
        <button type="button" onClick={() => append({ field: '', value: '' })}>
          Add Field
        </button>
      </form>
    </Modal>
  );
}

export default ManageFilterModal;
